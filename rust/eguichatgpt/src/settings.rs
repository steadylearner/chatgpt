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

        Settings {
            app_name: "ChatGPT Image Generator".to_string(),
            desktop_width: 300.0,
            desktop_height: 400.0,
            openai_api_key: "yours".to_string(),
            images_folder: "yours".to_string(), // Users/steadylearner/Desktop/images
        }
    }
}
