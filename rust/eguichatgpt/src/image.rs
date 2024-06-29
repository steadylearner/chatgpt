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
