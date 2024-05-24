// https://github.com/64bit/async-openai/blob/main/examples/assistants/src/main.rs

use async_openai::{
    types::{
        CreateAssistantRequestArgs, CreateMessageRequestArgs, CreateRunRequestArgs,
        CreateThreadRequestArgs, MessageContent, MessageRole, RunStatus,
    },
    Client,
};
use std::error::Error;
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
    }

    // once we have broken from the main loop we can delete the assistant and thread
    client.assistants().delete(assistant_id).await?;
    client.threads().delete(&thread.id).await?;

    Ok(())
}
