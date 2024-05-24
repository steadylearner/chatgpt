use dotenvy::dotenv;
use std::env;

lazy_static::lazy_static! {
    pub static ref SETTINGS: Settings = Settings::new();
}

pub struct Settings {
    pub openai_api_key: String,
    pub text_model: String,
    pub texts_folder: String,
    pub text_file_ext: String,
    pub images_folder: String,
    // pub image_model: String,
    // pub image_size: String,
    // pub quality: String,
}

impl Settings {
    pub fn new() -> Self {
        dotenv().expect(".env file not found");

        Settings {
            openai_api_key: env::var("OPENAI_API_KEY").expect("OPENAI_API_KEY must be set"),
            text_model: env::var("TEXT_MODEL").unwrap_or_else(|_| "gpt-3.5-turbo".to_string()), // gpt-4o
            texts_folder: env::var("TEXTS_FOLDER").unwrap_or_else(|_| "texts".to_string()),
            text_file_ext: env::var("TEXT_FILE_EXT").unwrap_or_else(|_| "md".to_string()),
            images_folder: env::var("IMAGES_FOLDER").unwrap_or_else(|_| "./images".to_string()),
            // image_model: env::var("IMAGE_MODEL").unwrap_or_else(|_| "dall-e-3".to_string()),
            // image_size: env::var("IMAGE_SIZE").unwrap_or_else(|_| "1024x1024".to_string()),
            // quality: env::var("QUALITY").unwrap_or_else(|_| "hd".to_string()),
        }
    }
}
