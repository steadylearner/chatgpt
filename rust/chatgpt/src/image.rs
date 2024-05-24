// https://github.com/64bit/async-openai/blob/main/examples/assistants/src/main.rs

use std::error::Error;
use std::time::Instant;

use dotenvy::dotenv;

use async_openai::{
    types::{
        CreateImageRequestArgs, Image, ImageModel, ImageQuality, ImageSize, ImagesResponse,
        ResponseFormat,
    },
    Client,
};

mod settings;
use settings::SETTINGS;

mod emojis;
use emojis::{BOT_EMOJI, USER_EMOJI};

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().expect(".env file not found");

    let client = Client::new();

    // let mut questioned_already = false;
    loop {
        // if questioned_already == false {
        //     println!(
        //         "{} Bot\nCan you describe the image you want to create with details?",
        //         &BOT_EMOJI
        //     );
        //     questioned_already = true;
        // } else {
        //     println!(
        //         "\n{} Bot\nCan you describe the image you want to create with details?",
        //         &BOT_EMOJI
        //     );
        // }
        println!(
            "\n{} Bot\nCan you describe the image you want to create with details?",
            &BOT_EMOJI
        );

        println!("\n{} You", &USER_EMOJI);
        let mut your_message = String::new();
        std::io::stdin().read_line(&mut your_message).unwrap();

        // break out of the loop if the user enters !quit
        if your_message.trim() == "!quit" {
            break;
        }

        let start_time = Instant::now();

        let request = CreateImageRequestArgs::default()
            .prompt(&your_message)
            .n(1)
            .response_format(ResponseFormat::Url)
            .model(ImageModel::DallE3) // Use this the result is incomparable.
            .size(ImageSize::S1024x1024)
            .quality(ImageQuality::HD)
            .user("async-openai") // Use another value if you want
            .build()?;

        let response: ImagesResponse = client.images().create(request).await?;

        let elapsed_time = start_time.elapsed();

        // It did take 17.89s to create the image.
        println!("\n{} Bot\nIt did take {:.2?} to create the image.", &BOT_EMOJI, elapsed_time);

        let chat_gpt_response = response.clone();

        let ImagesResponse { created: _, data } = response;

        let image: &Image = data[0].as_ref();

        match image {
            Image::Url {
                url,
                revised_prompt: _,
            } => {
                // cprint(
                //     f"""\n🤖 Bot\nHere is the link to the image.\n\n{image_url}\n""",
                // )
                println!(
                    "\n{} Bot\nHere is the link to the image.\n\n{}\n",
                    &BOT_EMOJI, url
                );
                // if let Some(prompt) = revised_prompt {
                //     println!("Revised Prompt: {}", prompt);
                // }
            }
            Image::B64Json {
                b64_json,
                revised_prompt: _,
            } => {
                println!("Base64 JSON: {}", b64_json);
                // if let Some(prompt) = revised_prompt {
                //     println!("Revised Prompt: {}", prompt);
                // }
            }
        }

        println!("{} Bot\nDo you want to save the image?", &BOT_EMOJI);
        let mut save_image = String::new();
        std::io::stdin().read_line(&mut save_image).unwrap();

        if save_image.to_lowercase().starts_with("y") {
            let paths = chat_gpt_response.save(&SETTINGS.images_folder).await?;
            paths.iter().for_each(|path| {
                // TODO
                // You will probably able to use
                // path.as_path() to read the file and use a name you want instead.

                println!(
                    "\n{} Bot\n Image file path is {}",
                    &BOT_EMOJI,
                    path.display()
                )
            });
            // .for_each(|path| println!("\n{} Bot\nHere is the link to the image.\n\n{}\n", &BOT_EMOJI, path.display()));
        }
    }

    Ok(())
}
