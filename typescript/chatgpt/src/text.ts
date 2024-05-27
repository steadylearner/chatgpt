import fs from "fs";
import path from "path";

import { OPENAI_BOT, botMessage, botQuestion } from "./bot";
import { BOT_EMOJI, USER_EMOJI } from "./constants/emojis";
import { TEXTS_FOLDER, TEXT_FILE_EXT } from "./settings";

async function createText() {
  let questionedAlready = false;
  while (true) {
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
