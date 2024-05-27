// import { OPENAI_API_KEY } from "./settings";

// console.log(OPENAI_API_KEY)

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
