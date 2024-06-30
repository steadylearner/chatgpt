<!-- # How to make Rust Desktop App with Egui and ChatGPT -->

[Rust website]: https://www.rust-lang.org
[ChatGPT]: https://chatgpt.com
[async-openai-rust]: https://github.com/64bit/async-openai
[openai-api-keys]: https://platform.openai.com/account/api-keys
[openai-api-limits]: https://platform.openai.com/settings/organization/limits 
[openai-api-billing]: https://platform.openai.com/settings/organization/billing
[openai-api-usage]: https://platform.openai.com/usage 

[cargo-bundle]: https://github.com/burtonageo/cargo-bundle

[Egui]: https://github.com/emilk/egui
[native_dialog]: https://github.com/native-dialog-rs/native-dialog-rs

[Steadylearner GitHub]: https://github.com/steadylearner
[Steadylearner ChatGPT repository]: https://github.com/steadylearner/chatgpt

[Hire me]: https://www.onlycoiners.com/user/steadylearner/product/i-will-help-you-to-build-a-full-stack-app-bot-or-smart-contr

[ChatGPT YouTube]: https://www.youtube.com/watch?v=JY7H286Jubg

In this post, you will learn how to make a desktop app with Rust [Egui] and ChatGPT with the [async-openai-rust] package. 

[You can find the code used here in this repository.][Steadylearner ChatGPT repository] It also has examples in other programming languages and blog posts for them.

[If you haven't used ChatGPT for images, you might find this useful.][ChatGPT YouTube]

I plan to include more examples with other programming languages, so feel free to like the repository if you find it useful and want to see more examples.

[You can also hire me if you need a 
full stack developer for your projects.][Hire me]

## Prerequisites

You can skip these and go to **Code** if you have done similar things already.

1. Install Rust
2. Set up dev environnment
3. Get OpenAI API key
4. Install Rust packages

### 1. Install Rust

If you don't already have Rust installed, visit the [Rust website] and follow the guide to download it. To check if Rust is installed, use your console and type this.

```console
$rustc --version
```

Depending on OS you use, you might see an instruction to install it.

[You can find more information about the installation process here.](https://doc.rust-lang.org/beta/book/ch01-01-installation.html#installation)

### 2. Set up dev environment

Create a Rust dev environment to manage your project. Here, we will name the project chatgpt, but you can choose any name.

```console
$cargo new eguichatgpt
```

Then, you will see these are created in your new chatgpt folder.

```console
Cargo.toml src/main.rs
```

Use `$cd chatgpt` and `$cargo run --release` to see if it worked.

You should see `Hello, world!` in your console.

Later, you can optionally use commands like these to help your development process.

```console
$cargo fmt
$cargo watch -x 'run --release'
```

### 3. Get OpenAI API key

Sign up on the OpenAI platform and get your API key from the [API keys page][openai-api-keys]. Store it securely as you will need it to access the OpenAI API later.

You can find more information about OpenAI API usage, billing, and limits at the following links.

- [API Limits][openai-api-limits]
- [API Billing][openai-api-billing]
- [API Usage][openai-api-usage]

### 4. Install Rust packages

You can install pacakges required for this post with Cargo.toml similar to this.

```toml
[package]
name = "eguichatgpt"
description = "An example ChatGPT Rust egui application by Steadylearner"
version = "0.1.0"
edition = "2021"

# https://crates.io/category_slugs
[package.metadata.bundle]
name = "ChatGPT"
icon = ["icon.png"]
version = "1.0.0"
copyright = "Copyright steadylearnerdev@gmail.com. All rights reserved."
category = "Developer Tool"
short_description = "An example ChatGPT Rust egui application by Steadylearner"
long_description = """
An example ChatGPT Rust egui application by Steadylearner
"""

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
eframe = "0.27.2"

egui = "0.27.2"
async-openai = "0.23.3"
dotenvy = "0.15.7"
lazy_static = "1.4.0"

env_logger = { version = "0.11.3", default-features = false, features = [
  "auto-color",
  "humantime",
] }
tokio = { version = "1", features = ["full"] }

egui_extras = { version = "0.27.2", features = ["default", "all_loaders"] }
image = { version = "0.25.1", features = ["jpeg", "png"] }
reqwest = { version = "0.12.5", features = ["blocking"] }
native-dialog = "0.7.0"
```

## Code

### Configuration

First, create a .env file in your chatgpt folder with the following text.

```console
OPENAI_API_KEY=<YOURS>
```

Then, create a settings.rs file to manage your settings.

```rs
use dotenvy::dotenv;
use std::env;

lazy_static::lazy_static! {
    pub static ref SETTINGS: Settings = Settings::new();
}

pub struct Settings {
    pub app_name: String,
    pub desktop_width: f32,
    pub desktop_height: f32,
    pub openai_api_key: String,
    pub images_folder: String,
}

impl Settings {
    pub fn new() -> Self {
        // Use .env to test
        // dotenv().expect(".env file not found");

        // Settings {
        //     app_name: "ChatGPT Image Generator".to_string(),
        //     desktop_width: 300.0,
        //     desktop_height: 400.0,
        //     openai_api_key: env::var("OPENAI_API_KEY").expect("OPENAI_API_KEY must be set"),
        //     images_folder: env::var("IMAGES_FOLDER").unwrap_or_else(|_| "images".to_string()),
        // }

        // Use this for production
        Settings {
            app_name: "ChatGPT Image Generator".to_string(),
            desktop_width: 300.0,
            desktop_height: 400.0,
            openai_api_key: "yours".to_string(),
            images_folder: "yours".to_string(), // Users/steadylearner/Desktop/images
        }
    }
}
```

You can use your own openai_api_key and images folder that you want to save images later.

### A Desktop app

First, we will create a main.rs that includes the most logics for a desktop app to work.

```rs
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

struct SteadylearnerChatGptApp {
    input: String,
    description: Arc<Mutex<Option<String>>>,
    image: Arc<Mutex<Option<String>>>,
    loading: Arc<Mutex<bool>>,
}

impl Default for SteadylearnerChatGptApp {
    fn default() -> Self {
        Self {
            input: "".to_owned(),
            description: Arc::new(Mutex::new(None)),
            image: Arc::new(Mutex::new(None)),
            loading: Arc::new(Mutex::new(false)),
        }
    }
}

impl eframe::App for SteadylearnerChatGptApp {
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

// $cargo install cargo-bundle
// $cargo bundle --release
// Edit Cargo.toml
// https://github.com/burtonageo/cargo-bundle
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
                Box::<SteadylearnerChatGptApp>::default()
            }),
        )
    })
}
```

Here, we first set the desktop app state with this. They are also corresponds to the app parts we will use for this app.

```rs
struct SteadylearnerChatGptApp {
    input: String,
    description: Arc<Mutex<Option<String>>>,
    image: Arc<Mutex<Option<String>>>,
    loading: Arc<Mutex<bool>>,
}

impl Default for SteadylearnerChatGptApp {
    fn default() -> Self {
        Self {
            input: "".to_owned(),
            description: Arc::new(Mutex::new(None)),
            image: Arc::new(Mutex::new(None)),
            loading: Arc::new(Mutex::new(false)),
        }
    }
}
```

`Arc<Mutex>` is used for some of them to update app state with some async code for openai. In your app without any async call, String or bool data type without it should work.

Then, we have the `main` function that starts the gui app.

```rs
use tokio::runtime::Runtime;

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
                Box::<SteadylearnerChatGptApp>::default()
            }),
        )
    })
}
```

`runtime` relevant things is what helps you to use async code inside egui app with `tokio::task::spawn` and `await`.

`egui_extras::install_image_loaders(&cc.egui_ctx)` part helps a egui desktop show the image you request from openai later.

Then, we have the `update` function that most of the app logcis are handled.

```rs
fn update(&mut self, ctx: &egui::Context, _frame: &mut eframe::Frame) {
    egui::CentralPanel::default().show(ctx, |ui| {
        // Increase the button text size
        let button_text_style = egui::TextStyle::Button;
        let mut text_style = ctx.style().text_styles.clone();
        text_style.insert(button_text_style, egui::FontId::proportional(16.0));
        ui.style_mut().text_styles = text_style;

        // Wrap the other parts with ui.vertical or ui.horizontal if you want some spacings among some elements.
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

                        // Here we request an image to openapi with a user input
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
                    
                    // Users can optionally save images with download button.
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
```

Here, we handle show some error message when there was some issues or user input is empty with description.

When user use a proper prompt, the app requests an image url with this part.

```rs
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
```

We can use this becausee we included `runtime.block_on` part in the main function before and we can update the app state with an async function `fetch_image_url` from image.rs.

```rs
use std::error::Error;

use async_openai::{
    config::OpenAIConfig,
    types::{
        CreateImageRequestArgs, Image, ImageModel, ImageQuality, ImageSize, ImagesResponse,
        ResponseFormat,
    },
    Client,
};

use crate::settings::SETTINGS;

pub async fn fetch_image_url(input: &str) -> Result<String, Box<dyn Error>> {
    let config = OpenAIConfig::new().with_api_key(SETTINGS.openai_api_key.clone());

    let client = Client::with_config(config);

    let request = CreateImageRequestArgs::default()
        .prompt(input)
        .n(1)
        .response_format(ResponseFormat::Url)
        .model(ImageModel::DallE3) // Use this the result is incomparable.
        .size(ImageSize::S1792x1024)
        .quality(ImageQuality::HD)
        .user("async-openai") // Use another value if you want
        .build()?;

    let response: ImagesResponse = client.images().create(request).await?;

    // println!("{:#?}", &response);

    let ImagesResponse { created: _, data } = response;

    let image: &Image = data[0].as_ref();

    match image {
        Image::Url {
            url,
            revised_prompt: _,
        } => Ok(url.to_string()),
        _ => Err("Unexpected image output".into()),
    }
}
```

The users can also optionally download images from openai with download button.

```rs
if ui.button("Download").clicked() {
    if let Err(error) = download_and_save_file(image) {
        let mut description_guard = self.description.lock().unwrap();
        *description_guard = Some(error.to_string());
    } 
}
```

You can make `download_and_save_file` function similar to this. You can refer to [native_dialog] package docs.

It will first ask a user where to save an iamge and a file name to use. Then, it will use the url to download and save it with the information.

```rs
use native_dialog::FileDialog;
use std::fs::File;
use std::io::copy;

use reqwest;

use crate::settings::SETTINGS;

pub fn download_and_save_file(url: &str) -> Result<(), Box<dyn std::error::Error>> {
    let result = FileDialog::new()
        .set_location(&SETTINGS.images_folder)
        .add_filter("PNG Image", &["png"])
        .show_save_single_file()?; 

    if let Some(file_path) = result {
        let response = reqwest::blocking::get(url)?;
        let mut file = File::create(file_path)?;

        copy(&mut response.bytes().unwrap().as_ref(), &mut file)?;
    } 

    Ok(())
}
```

Then, you can use `$cargo run --release` or and 

## Packaging for distribution

With Rust [cargo-bundle], you can easily package a Rust app for distribution with these commands.

```console
$cargo install cargo-bundle
$cargo bundle --release
```

Before that, you will first have to edit your Cargo.toml to include more fields to help this process.

You can refer to this and edit yours.

``toml
[package]
name = "eguichatgpt"
description = "An example ChatGPT Rust egui application by Steadylearner"
version = "0.1.0"
edition = "2021"

# https://crates.io/category_slugs
[package.metadata.bundle]
name = "ChatGPT"
# name = "SteadylearnerChatGPT"
icon = ["icon.png"]
version = "1.0.0"
copyright = "Copyright steadylearnerdev@gmail.com. All rights reserved."
category = "Developer Tool"
short_description = "An example ChatGPT Rust egui application by Steadylearner"
long_description = """
An example ChatGPT Rust egui application by Steadylearner
"""
``````

Then, it will create a desktop app `eguichatgpt` in your target/release/bundle/osx/ SteadylearnerChatGPT.app/Contents/MacOS folder.

## Conclusion

You have learned how to create a simple desktop app with [Egui] and [async-openai-rust]. 

If you thought this post helpful, please share it and like the [ChatGPT repository][Steadylearner ChatGPT repository] and blog post.

You can also join [Rust community](https://www.onlycoiners.com/community/rust) with images created from the app described in this post. 

[You can also hire me if you need a full stack developer for your projects.][Hire me]