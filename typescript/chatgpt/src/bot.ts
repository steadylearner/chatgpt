import OpenAI from "openai";
import readlineSync from "readline-sync";

import { OPENAI_API_KEY } from "./settings";
import { BOT_EMOJI, USER_EMOJI } from "./constants/emojis";

const OPENAI_BOT = new OpenAI({
  apiKey: OPENAI_API_KEY,
});

function botQuestion(message: string): string {
  const userResponse = readlineSync.question(message);
  return userResponse;
}

function userMessage(text: string) {
  console.log(`
${USER_EMOJI} Bot 
${text}
  `);
}

function botMessage(text: string) {
  console.log(`
${BOT_EMOJI} Bot 
${text}`);
}

export { OPENAI_BOT, botQuestion, userMessage, botMessage };
