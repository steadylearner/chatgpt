import dotenv from "dotenv";
dotenv.config();

export const OPENAI_API_KEY = process.env.OPENAI_API_KEY;

export const TEXTS_FOLDER = "texts";
export const TEXT_FILE_EXT = "md";

export const IMAGES_FOLDER = "images";
export const IMAGE_SIZE = "1024x1024";
export const QUALITY = "hd"; // standard, hd
