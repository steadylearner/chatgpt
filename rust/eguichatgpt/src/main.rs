use eframe::egui;
use std::sync::{Arc, Mutex};
use tokio::runtime::Runtime;

mod settings;
use settings::SETTINGS;

mod consts;
use consts::TYPE_SOMETHING;

mod image;
use image::fetch_image_url;

mod file;
use file::download_and_save_file;

struct MyApp {
    input: String,
    description: Arc<Mutex<Option<String>>>,
    image: Arc<Mutex<Option<String>>>,
    loading: Arc<Mutex<bool>>,
}

impl Default for MyApp {
    fn default() -> Self {
        Self {
            input: "".to_owned(),
            description: Arc::new(Mutex::new(None)),
            image: Arc::new(Mutex::new(None)),
            loading: Arc::new(Mutex::new(false)),
        }
    }
}

impl eframe::App for MyApp {
    fn update(&mut self, ctx: &egui::Context, _frame: &mut eframe::Frame) {
        egui::CentralPanel::default().show(ctx, |ui| {
            let button_text_style = egui::TextStyle::Button;
            let mut text_style = ctx.style().text_styles.clone();
            text_style.insert(button_text_style, egui::FontId::proportional(16.0));
            ui.style_mut().text_styles = text_style;

            ui.vertical(|ui| {

                let mut description_guard = self.description.lock().unwrap();
                ui.vertical(|ui| {
                    ui.spacing_mut().item_spacing.y = 10.0;

                    ui.heading("Describe the image with details");

                    ui.text_edit_multiline(&mut self.input);

                    if let Some(description) = &*description_guard {
                        ui.add(egui::Label::new(description));
                    }
 
                });

                ui.horizontal(|ui| {
                    ui.set_height(36.0);

                    if ui.button("Submit").clicked() {
                        if self.input.is_empty() {
                            *description_guard = Some(TYPE_SOMETHING.to_owned());
                        } else {
                            *description_guard = None;

                            let mut image_guard = self.image.lock().unwrap();
                            *image_guard = None;

                            let mut loading_guard = self.loading.lock().unwrap();
                            *loading_guard = true;

                            let input = self.input.clone();
                            let image_mutex = self.image.clone();
                            let description_mutex = self.description.clone();
                            let loading_mutex = self.loading.clone();

                            tokio::task::spawn(async move {
                                match fetch_image_url(&input).await {
                                    Ok(url) => {
                                        let mut image_guard = image_mutex.lock().unwrap();
                                        *image_guard = Some(url);
                                    }
                                    Err(error) => {
                                        let mut description_guard =
                                            description_mutex.lock().unwrap();
                                        *description_guard = Some(error.to_string());
                                    }
                                }

                                let mut loading_guard = loading_mutex.lock().unwrap();
                                *loading_guard = false;
                            });
                        }
                    }

                    if ui.button("Reset").clicked() {
                        self.input = "".to_owned();

                        let image_mutex = self.image.clone();
                        let description_mutex = self.description.clone();
                        let loading_mutex = self.loading.clone();

                        tokio::task::spawn(async move {
                            let mut image_guard = image_mutex.lock().unwrap();
                            *image_guard = None;
                            let mut description_guard = description_mutex.lock().unwrap();
                            *description_guard = None;

                            let mut loading_guard = loading_mutex.lock().unwrap();
                            *loading_guard = false;
                        });
                    }
                });

                let loading_guard = self.loading.lock().unwrap();
                if *loading_guard {
                    ui.spinner();
                }

                let image_guard = self.image.lock().unwrap();
                if let Some(image) = &*image_guard {
                    ui.vertical(|ui| {
                        ui.spacing_mut().item_spacing.y = 10.0;
                        
                        if ui.button("Download").clicked() {
                            if let Err(error) = download_and_save_file(image) {
                                let mut description_guard = self.description.lock().unwrap();
                                *description_guard = Some(error.to_string());
                            } 
                        }

                        let image_widget = egui::Image::new(image).rounding(10.0).max_width(300.0);
                        ui.add(image_widget);
                    });
                }
            });
        });
    }
}

// TODO
// 1. Find how to make a production app with it

// $cargo new eguichatgpt
// $cargo watch -x 'run --release'
fn main() -> eframe::Result<(), eframe::Error> {
    env_logger::init(); 

    let runtime = Runtime::new().unwrap();

    // Configure the application window and options
    let options = eframe::NativeOptions {
        viewport: egui::ViewportBuilder::default()
            .with_inner_size([SETTINGS.desktop_width, SETTINGS.desktop_height]),
        ..Default::default()
    };

    // Run the Egui application within Tokio runtime
    runtime.block_on(async {
        eframe::run_native(
            &SETTINGS.app_name,
            options,
            Box::new(|cc| {
                egui_extras::install_image_loaders(&cc.egui_ctx);
                Box::<MyApp>::default()
            }),
        )
    })
}
