<!-- # How to use ChatGPT with Rust -->

[Rust website]: https://www.rust-lang.org
[ChatGPT]: https://chatgpt.com
[async-openai-rust]: https://github.com/64bit/async-openai
[openai-api-keys]: https://platform.openai.com/account/api-keys
[openai-api-limits]: https://platform.openai.com/settings/organization/limits 
[openai-api-billing]: https://platform.openai.com/settings/organization/billing
[openai-api-usage]: https://platform.openai.com/usage 

[Steadylearner GitHub]: https://github.com/steadylearner
[Steadylearner ChatGPT repository]: https://github.com/steadylearner/chatgpt

[Hire me]: https://www.onlycoiners.com/user/steadylearner/product/i-will-help-you-to-build-a-full-stack-app-bot-or-smart-contr

[ChatGPT YouTube]: https://www.youtube.com/watch?v=JY7H286Jubg

In this post, you will learn how to use ChatGPT with Rust using the official [async-openai-rust] package. 

[You can find the code used here in this repository.][Steadylearner ChatGPT repository] It also has examples in other programming languages and blog posts for them.

[If you haven't used ChatGPT for images, you might find this useful.][ChatGPT YouTube]

I plan to include more examples with other programming languages, so feel free to like the repository if you find it useful and want to see more examples.

[You can also hire me if you need a full stack developer for your projects.][Hire me]

## Prerequisites

You can skip these and go to **Code** if you have done similar things already.

1. Install Rust
2. Set up dev envrionment
3. Get OpenAI API key
4. Install Rust packages

### 1. Install Rust

If you don't already have Rust installed, visit the [Rust website] and follow the guide to download it. To check if Rust is installed, use your console and type this.

```console
$rustc --version
```

Depending on OS you use, you might see an instruction to install it.

[You can find more information about the installation process here.](https://doc.rust-lang.org/beta/book/ch01-01-installation.html#installation)

### 2. Set up virtual environment

Create a Rust virtual environment to manage your dependencies. Here, we will name the environment chatgpt, but you can choose any name

```console
$cargo new chatgpt
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

You can install pacakges required for this post with this.

```console
$cargo add dotenvy async-openai tokio
```

You can also use `$cargo add tracing-subscriber` for examples from [async-openai-rust] package.

You can also use Cargo.toml similar to this.

```toml
[package]
name = "chatgpt"
version = "0.1.0"
edition = "2021"

[dependencies]
async-openai = "0.21.0"
dotenvy = "0.15.7"
lazy_static = "1.4.0"
tokio = { version = "1.37.0", features = ["full", "rt-multi-thread"] }
tracing-subscriber = { version = "0.3.18", features = ["env-filter"] }

# $cargo run --bin text
[[bin]]
name = "text"
path = "src/text.rs"

# $cargo run --bin image
[[bin]]
name = "image"
path = "src/image.rs"
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
    pub openai_api_key: String,
    pub text_model: String,
    pub texts_folder: String,
    pub text_file_ext: String,
    pub images_folder: String,
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
        }
    }
}
```

Finally, we will see two examples that you might find useful and make some variations later.

### A CLI chatbot

First, we will see a simple code snippet for a CLI chatbot.

You can save the text below as **text.rs** file.

```rs
use std::error::Error;
use std::time::Instant;

use async_openai::{
    types::{
        CreateAssistantRequestArgs, CreateMessageRequestArgs, CreateRunRequestArgs,
        CreateThreadRequestArgs, MessageContent, MessageRole, RunStatus,
    },
    Client,
};
use tracing_subscriber::{fmt, layer::SubscriberExt, util::SubscriberInitExt, EnvFilter};

use dotenvy::dotenv;

mod settings;
use settings::SETTINGS;

mod emojis;
use emojis::{BOT_EMOJI, USER_EMOJI};

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().expect(".env file not found");

    std::env::set_var("RUST_LOG", "ERROR");

    // Setup tracing subscriber so that library can log the errors
    tracing_subscriber::registry()
        .with(fmt::layer())
        .with(EnvFilter::from_default_env())
        .init();

    let query = [("limit", "1")]; //limit the list responses to 1 message

    let client = Client::new();

    let thread_request = CreateThreadRequestArgs::default().build()?;
    let thread = client.threads().create(thread_request.clone()).await?;

    // ChatGPT FAQ Rust CLI bot
    println!("{} Bot\nWhat is the name of your assistant?", &BOT_EMOJI);
    let mut assistant_name = String::new();
    std::io::stdin().read_line(&mut assistant_name).unwrap();

    // You should respond professionally
    println!(
        "\n{} Bot\nWhat is the instruction set for your new assistant?",
        &BOT_EMOJI
    );
    let mut instructions = String::new();
    std::io::stdin().read_line(&mut instructions).unwrap();

    let assistant_request = CreateAssistantRequestArgs::default()
        .name(&assistant_name)
        .instructions(&instructions)
        .model(&SETTINGS.text_model)
        .build()?;
    let assistant = client.assistants().create(assistant_request).await?;
    let assistant_id = &assistant.id;

    let mut questioned_already = false;
    loop {
        // 1.
        if questioned_already == false {
            println!("\n{} Bot\nHow can I help you?", &BOT_EMOJI);
            questioned_already = true;
        } else {
            println!("{} Bot\nAnything else I can help?", &BOT_EMOJI);
        }
        // "What does steadylearner mean?"
        println!("\n{} You", &USER_EMOJI);
        let mut your_message = String::new();
        std::io::stdin().read_line(&mut your_message).unwrap();

        // break out of the loop if the user enters !quit
        if your_message.trim() == "!quit" {
            break;
        }

        let start_time = Instant::now();

        // create a message for the thread
        let message = CreateMessageRequestArgs::default()
            .role(MessageRole::User)
            .content(your_message.clone())
            .build()?;

        // attach message to the thread
        let _message_obj = client
            .threads()
            .messages(&thread.id)
            .create(message)
            .await?;

        // create a run for the thread
        let run_request = CreateRunRequestArgs::default()
            .assistant_id(assistant_id)
            .build()?;
        let run = client
            .threads()
            .runs(&thread.id)
            .create(run_request)
            .await?;

        // wait for the run to complete
        let mut awaiting_response = true;
        while awaiting_response {
            // retrieve the run
            let run = client.threads().runs(&thread.id).retrieve(&run.id).await?;
            // check the status of the run
            match run.status {
                RunStatus::Completed => {
                    awaiting_response = false;
                    // once the run is completed we
                    // get the response from the run
                    // which will be the first message
                    // in the thread

                    // retrieve the response from the run
                    let response = client.threads().messages(&thread.id).list(&query).await?;
                    // get the message id from the response
                    let message_id = response.data.get(0).unwrap().id.clone();
                    // get the message from the response
                    let message = client
                        .threads()
                        .messages(&thread.id)
                        .retrieve(&message_id)
                        .await?;
                    // get the content from the message
                    let content = message.content.get(0).unwrap();
                    // get the text from the content
                    let text = match content {
                        MessageContent::Text(text) => text.text.value.clone(),
                        MessageContent::ImageFile(_) => {
                            panic!("images are not supported in the terminal")
                        }
                    };
                    println!("\n{} Bot\n{}\n", &BOT_EMOJI, text);

                    // TODO
                    // Include a question if a user want to save the text or not
                }
                RunStatus::Failed => {
                    awaiting_response = false;
                    println!("\n{} Bot\n Failed: {:#?}", &BOT_EMOJI, run);
                }
                RunStatus::Queued => {
                    println!("\n{} Bot\n Run Queued", &BOT_EMOJI);
                }
                RunStatus::Cancelling => {
                    println!("\n{} Bot\nRun Cancelling", &BOT_EMOJI);
                }
                RunStatus::Cancelled => {
                    println!("\n{} Bot\nRun Cancelled", &BOT_EMOJI);
                }
                RunStatus::Expired => {
                    println!("\n{} Bot\nRun Expired", &BOT_EMOJI);
                }
                RunStatus::RequiresAction => {
                    println!("\n{} Bot\nRun Requires Action", &BOT_EMOJI);
                }
                RunStatus::InProgress => {
                    println!("\n{} Bot\nWaiting for response...", &BOT_EMOJI);
                }
            }
            //wait for 1 second before checking the status again
            tokio::time::sleep(tokio::time::Duration::from_secs(1)).await;
        }

        let elapsed_time = start_time.elapsed();
        println!("\n{} Bot\nIt did take {:.2?} to create the response.", &BOT_EMOJI, elapsed_time);
    }

    // once we have broken from the main loop we can delete the assistant and thread
    client.assistants().delete(assistant_id).await?;
    client.threads().delete(&thread.id).await?;

    Ok(())
}
```

We already included this code snippet to Cargo.toml file.

```toml
[[bin]]
name = "text"
path = "src/text.rs"
```

So we can test it with `$cargo run --bin text` and you will see some messages like this.

You can use your own question instead of "What does steadylearner mean?"

```console
🤖 Bot 
How can I help you?

🧑 You
What does steadylearner mean?

🤖 Bot
Steadylearner likely means someone who is constantly learning and seeking knowledge in a consistent and steady manner. This person is dedicated to their personal growth and development through continuous learning and improvement.

🤖 Bot 
Anything else I can help?
```

So how does it work? There are a few important parts but we will see only important ones.

1. We want to ask questions to the bot until we stop the bot with `Ctrl + C`, `!quit` or others that results in SIGINT signal.
2. The time for the ChatGPT API to respond varies depending on your question and it might take very long so we will have some feedbacks to decrease response time by measuring how much it takes.

So that was the simple code snippet needed to respond to your questions.

The **gpt-3.5-turbo** option was used here to save resources but you can also use **gpt-4o** option.

Not so interesting because you can use more features at [ChatGPT] website. Therefore, we will make something a little bit more useful variation.

[You can see the original code example used as a reference by the package owner is here.](https://github.com/64bit/async-openai/blob/main/examples/assistants/src/main.rs)

You can also test that.

### A CLI image creator

The code snippet used here is almost the same to the previous one but we will use **dall-e-3** model to make images. You can use it for profile image, blog cover, logo etc.

```rs
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

    loop {
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

        println!("\n{} Bot\nIt did take {:.2?} to create the image.", &BOT_EMOJI, elapsed_time);

        let chat_gpt_response = response.clone();

        let ImagesResponse { created: _, data } = response;

        let image: &Image = data[0].as_ref();

        match image {
            Image::Url {
                url,
                revised_prompt: _,
            } => {
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
        }
    }

    Ok(())
}
```

You can test it with `$cargo run --bin image` command. Then, describe the image you want to create with details.

You will see similar messages you read in the previous process but you will see an image created with url and how much time it did take.

You can use another message instead of **Can you make a cover image for blog post with title "How to use chatgpt with Rust"** used here.

```console
🤖 Bot 
Can you describe the image you want to create with details?

🧑 You
Can you make a cover image for blog post with title "How to use chatgpt with Rust"

🤖 Bot
Here is the link to the image.

https://oaidalleapiprodscus.blob.core.windows.net/private

Do you want to save the image?
y
```

The profile and the cover image on this blog was created with it. 

You can also make your own images.

## Conclusion

You have learned how to create simple CLIs for text and image using Rust and [async-openai-rust]. 

You can extend these examples to create more complex applications. 

If you thought this post helpful, please share it and like the [ChatGPT repository][Steadylearner ChatGPT repository] and blog post.

[You can also hire me if you need a full stack developer for your projects.][Hire me]