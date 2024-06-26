// $cargo run --release
// $cargo watch -x 'run --release'

use dotenvy::dotenv;
use std::env;

use async_openai::{
    types::{CreateImageRequestArgs, ImageSize, ResponseFormat},
    Client,
};
use std::error::Error;

// fn main() {
//     dotenv().expect(".env file not found");
//     let _open_api_key = env::var("OPENAI_API_KEY").expect("OPENAI_API_KEY must be set");
//     // println!("{}", _open_api_key);
// }

// This works
#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().expect(".env file not found"); // $export OPENAI_API_KEY='sk-...'

    // let _open_api_key = env::var("OPENAI_API_KEY").expect("OPENAI_API_KEY must be set");
    // // println!("{}", _open_api_key);

    // create client, reads OPENAI_API_KEY environment variable for API key.

    let client = Client::new();

    let request = CreateImageRequestArgs::default()
        .prompt("cats on sofa and carpet in living room")
        .n(1)
        .response_format(ResponseFormat::Url)
        .size(ImageSize::S256x256)
        .user("async-openai")
        .build()?;

    let response = client.images().create(request).await?;

    // Download and save images to ./data directory.
    // Each url is downloaded and saved in dedicated Tokio task.
    // Directory is created if it doesn't exist.
    let paths = response.save("./data").await?;

    paths
        .iter()
        .for_each(|path| println!("Image file path: {}", path.display()));

    Ok(())
}
