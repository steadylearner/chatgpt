<!-- # How to use ChatGPT with Rust -->

[Rust website]: https://www.rust-lang.org
[ChatGPT]: https://chatgpt.com
[openai-node]: https://github.com/openai/openai-node
[openai-api-keys]: https://platform.openai.com/account/api-keys
[openai-api-limits]: https://platform.openai.com/settings/organization/limits 
[openai-api-billing]: https://platform.openai.com/settings/organization/billing
[openai-api-usage]: https://platform.openai.com/usage 

[Steadylearner GitHub]: https://github.com/steadylearner
[Steadylearner ChatGPT repository]: https://github.com/steadylearner/chatgpt

[Hire me]: https://www.onlycoiners.com/user/steadylearner/product/i-will-help-you-to-build-a-full-stack-app-bot-or-smart-contr

[ChatGPT YouTube]: https://www.youtube.com/watch?v=JY7H286Jubg

In this post, you will learn how to use ChatGPT with TypeScript using the official [openai-node] package. 

[You can find the code used here in this repository.][Steadylearner ChatGPT repository] It also has examples in other programming languages and blog posts for them.

[If you haven't used ChatGPT for images, you might find this useful.][ChatGPT YouTube]

I plan to include more examples with other programming languages, so feel free to like the repository if you find it useful and want to see more examples.

[You can also hire me if you need a full stack developer for your projects.][Hire me]

## Prerequisites

You can skip these and go to **Code** if you have done similar things already.

1. Install JavaScript
2. Set up dev envrionment
3. Get OpenAI API key
4. Install required packages

### 1. Install Node

We will use TypeScript here but to use it you have to install Node first. [Visit this page for NVM](https://github.com/nvm-sh/nvm) and follow the guide to download it. To check if Node is installed, use your console and type this.

```console
$node --version
```

### 2. Set up dev environment

Create a Node js dev environment to manage your dependencies. Here, we will name the project chatgpt, but you can choose any name.

```console
$yarn init
```

Then, respond some questions and you will see these are created in your new chatgpt folder.

```console
package.json
```

[You can include more files similar what you can see here.](https://github.com/steadylearner/chatgpt/tree/main/typescript/chatgpt)

### 3. Get OpenAI API key

Sign up on the OpenAI platform and get your API key from the [API keys page][openai-api-keys]. Store it securely as you will need it to access the OpenAI API later.

You can find more information about OpenAI API usage, billing, and limits at the following links.

- [API Limits][openai-api-limits]
- [API Billing][openai-api-billing]
- [API Usage][openai-api-usage]

### 4. Install required packages

You can install pacakges required for this post with `$yarn` and **package.json** similar to this.

```json
{
  "name": "chatgpt",
  "version": "1.0.0",
  "author": "Steadylearner @steadylearner steadylearnerdev@email.com",
  "description": "",
  "main": "src/index.ts",
  "scripts": {
    "dev": "npx nodemon",
    "build": "tsc",
    "test": "npx tsx ./src/index.ts",
    "test:deploy": "node build/index.js",
    "lint": "prettier -w ."
  },
  "dependencies": {
    "@types/node": "^18.15.0",
    "@types/node-fetch": "^2.6.2",
    "axios": "^1.7.2",
    "commander": "^12.1.0",
    "dotenv": "^16.0.3",
    "openai": "^4.47.1",
    "readline-sync": "^1.4.10",
    "ts-node": "^10.9.1",
    "typescript": "^4.9.5"
  },
  "devDependencies": {
    "prettier": "2.7.1"
  }
}
```

It will be easier if you just clone the repository and start from there.

```console
$git clone https://github.com/steadylearner/chatgpt.git
```

## Code

### Configuration

First, create a .env file in your chatgpt folder with the following text.

```console
OPENAI_API_KEY=<YOURS>
```

Then, create a settings.ts file to manage your settings.

```ts
import dotenv from "dotenv";
dotenv.config();

export const OPENAI_API_KEY = process.env.OPENAI_API_KEY;

export const TEXTS_FOLDER = "texts";
export const TEXT_FILE_EXT = "md";

export const IMAGES_FOLDER = "images";
export const IMAGE_SIZE = "1024x1024";
export const QUALITY = "hd"; // standard, hd
```

Finally, we will see two examples that you might find useful and make some variations later.

But before we code on, we will write **index.ts** file similar to this to help you test the code snippets we will have later.

```ts
require("dotenv").config();
const { Command } = require("commander");
const { createText } = require("./text.ts");
const { createImage } = require("./image.ts");

const program = new Command();

program
  .name("ChatGPT")
  .version("0.0.1")
  .description("CLI tool for AI text and image");

// $yarn test text
program
  .command("text")
  .description("Create text based on a prompt")
  .action(async () => {
    await createText();
  });

// $yarn test image
program
  .command("image")
  .description("Create an image based on a prompt")
  .action(async () => {
    await createImage();
  });

program.parse(process.argv);
```

### A CLI chatbot

First, we will see a simple code snippet for a CLI chatbot.

You can save the text below as **text.ts** file.

```ts
import fs from "fs";
import path from "path";

import { OPENAI_BOT, botMessage, botQuestion } from "./bot";
import { BOT_EMOJI, USER_EMOJI } from "./constants/emojis";
import { TEXTS_FOLDER, TEXT_FILE_EXT } from "./settings";

async function createText() {
  let questionedAlready = false;
  while (true) {
    // 1.
    if (questionedAlready === false) {
      botMessage("How can I help you?");
      questionedAlready = true;
    } else {
      botMessage("Anything else I can help?");
    }

    console.log(`\n${USER_EMOJI} You`);
    const yourMessage = botQuestion(""); // What does steadylearner mean?

    if (yourMessage === "!quit") {
      break
    }

    if (yourMessage.length > 0) {
      // 2.   
      const startTime = Date.now();

      const { choices } = await OPENAI_BOT.chat.completions.create({
        messages: [{ role: "user", content: yourMessage }],
        model: "gpt-3.5-turbo",
      });
      const text = choices[0].message.content;
      botMessage(text);

      const endTime = Date.now();

      const elapsedTime = (endTime - startTime) / 1000;
      botMessage(
        `It took ${elapsedTime.toFixed(2)} seconds to create the response.\n`
      );

      const saveText = botQuestion(
        `${BOT_EMOJI} Do you want to save it?\n`
      );
      if (saveText.toLowerCase().startsWith("y")) {
        if (!fs.existsSync(TEXTS_FOLDER)) {
          fs.mkdirSync(TEXTS_FOLDER);
        }

        let textFileName = await botQuestion(
          `\n${BOT_EMOJI}What is the name of the text file?\n`
        );
        if (textFileName === "") {
          const currentTimestamp = Math.floor(Date.now() / 1000);
          textFileName = currentTimestamp.toString();
        }

        const textFilename = `${textFileName}.${TEXT_FILE_EXT}`;
        const textFilePath = path.join(TEXTS_FOLDER, textFilename);

        fs.writeFileSync(textFilePath, text);
        console.log(
          `\nThe response ${textFileName} was saved to ${textFilePath}`
        );
      }
    }
  }
}

export { createText };
```

We already included this code snippet to index.ts file.

```ts
program
  .command("text")
  .description("Create text based on a prompt")
  .action(async () => {
    await createText();
  });
```

So we can test it with `$yarn test text` and you will see some messages like this.

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

### A CLI image creator

The code snippet used here is almost the same to the previous one but we will use **dall-e-3** model to make images. You can use it for profile image, blog cover, logo etc.

```ts
import fs from "fs";
import path from "path";

import { OPENAI_BOT, botMessage, botQuestion } from "./bot";
import { BOT_EMOJI, USER_EMOJI } from "./constants/emojis";
import { IMAGES_FOLDER, IMAGE_SIZE, QUALITY, TEXTS_FOLDER, TEXT_FILE_EXT } from "./settings";
import axios from "axios";

async function createImage() {
  while (true) {
    // Steadylearner
    botMessage("Can you describe the image you want to create with details?");

    console.log(`\n${USER_EMOJI} You`);
    const yourMessage = botQuestion(""); // What does steadylearner mean?

    if (yourMessage === "!quit") {
      break
    }

    if (yourMessage.length > 0) {
      const startTime = Date.now();

      const chatGptResponse = await OPENAI_BOT.images.generate({
        model: "dall-e-3",
        prompt: yourMessage,
        size: IMAGE_SIZE,
        quality: QUALITY,
        n: 1,
      });
      const imageUrl = chatGptResponse.data[0].url;

      const endTime = Date.now();

      botMessage(
        `Here is the link to the image.\n\n${imageUrl}`,
      )

      const elapsedTime = (endTime - startTime) / 1000;
      botMessage(
        `It took ${elapsedTime.toFixed(2)} seconds to create the image.\n`
      );

      const saveText = botQuestion(
        `${BOT_EMOJI} Do you want to save it?\n`
      );
      if (saveText.toLowerCase().startsWith("y")) {
        if (!fs.existsSync(IMAGES_FOLDER)) {
          fs.mkdirSync(IMAGES_FOLDER);
        }

        let imageFileName = await botQuestion(
          `\n${BOT_EMOJI}What is the name of the image?\n`
        );
        if (imageFileName === "") {
          const currentTimestamp = Math.floor(Date.now() / 1000);
          imageFileName = currentTimestamp.toString();
        }

        const imageFilename = `${imageFileName}.png`;
        const imageFilePath = path.join(IMAGES_FOLDER, imageFilename);

        try {
          const response = await axios.get(imageUrl, { responseType: 'arraybuffer' });
          if (response.status === 200) {
            try {
              fs.writeFileSync(imageFilePath, response.data, 'binary')
              console.log(`\nThe image ${imageFilename} was saved to ${imageFilePath}`);
            } catch (error) {
              botMessage(`Unable to save the image with error below. \n\n ${error}`);
            }
          } else {
            botMessage("Unable to save the image");
          }
        } catch (error) {
          botMessage(`Unable to save the image with error below. \n\n ${error}`);
        }
      }
    }
  }
}

export { createImage };

```

You can test it with `$yarn test image` command. Then, describe the image you want to create with details.

You will see similar messages you read in the previous process but you will see an image created with url and how much time it did take.

You can use another message instead of **Can you make a cover image for blog post with title "How to use chatgpt with TypeScript"** used here.

```console
🤖 Bot 
Can you describe the image you want to create with details?

🧑 You
Can you make a cover image for blog post with title "How to use chatgpt with TypeScript"

🤖 Bot
Here is the link to the image.

https://oaidalleapiprodscus.blob.core.windows.net/private

Do you want to save the image?
y
```

The profile and the cover image on this blog was created with it. 

You can also make your own images.

## Conclusion

You have learned how to create simple CLIs for text and image using Rust and [openai-node]. 

You can extend these examples to create more complex applications. 

If you thought this post helpful, please share it and like the [ChatGPT repository][Steadylearner ChatGPT repository] and blog post.

[You can also hire me if you need a full stack developer for your projects.][Hire me]