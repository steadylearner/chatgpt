from dotenv import load_dotenv
from pathlib import Path
import os

env_path = Path(".") / ".env"
load_dotenv(dotenv_path=env_path)

OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")

TEXTS_FOLDER = "texts"
TEXT_FILE_EXT = "md"

IMAGES_FOLDER = "images"
IMAGE_SIZE = "1024x1024"
QUALITY = "hd" # standard, hd