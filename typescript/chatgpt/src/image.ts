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
      const endTime = Date.now();

      const imageUrl = chatGptResponse.data[0].url;

      botMessage(
        `Here is the link to the image.\n\n${imageUrl}`,
      )

      const elapsedTime = (endTime - startTime) / 1000;
      botMessage(
        `It took ${elapsedTime.toFixed(2)} seconds to create the image.\n`
      );

      const saveImage = botQuestion(
        `${BOT_EMOJI} Do you want to save it?\n`
      );
      if (saveImage.toLowerCase().startsWith("y")) {
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
