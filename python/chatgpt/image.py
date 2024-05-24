# $pip install black $black -q .

import os
import requests
import time

from openai import OpenAI
from settings import (
  OPENAI_API_KEY, 

  IMAGES_FOLDER,
  IMAGE_SIZE, 
  QUALITY,
)

from termcolor import cprint

chat_gpt_model = "dall-e-3" 

client = OpenAI(
    # This is the default and can be omitted
    api_key=OPENAI_API_KEY,
)

bot_message = lambda text: cprint(
    f"""
🤖 Bot 
{text}
""",
    # color="yellow",
    # attrs=["bold"]
)

while True:
    bot_message("Can you describe the image you want to create with details?")

    # I want a cyber punk cat with red bublue gum popping
    print("🧑 You")
    your_message = input()

    # Start the timer
    start_time = time.time()

    # TODO
    # Include the timer here.
    chat_gpt_response = client.images.generate(
        model=chat_gpt_model,
        prompt=your_message,
        size=IMAGE_SIZE,
        quality=QUALITY,
        n=1,
    )

    # Stop the timer
    end_time = time.time()

    # Calculate the time difference
    elapsed_time = end_time - start_time

    # print("response")
    # print(response)
    created = chat_gpt_response.created
    revised_prompt = chat_gpt_response.data[0].revised_prompt
    image_url = chat_gpt_response.data[0].url

    # cprint(
    #     f"""\n🤖 Bot\nHere is the revisted prompt of your message.\n\n{revised_prompt}""",
    #     # color="yellow", attrs=["bold"]
    # )

    cprint(
        f"""\n🤖 Bot\nHere is the link to the image.\n\n{image_url}\n""",
        # color="yellow", attrs=["bold"]
    )

    print(f"\n🤖It did take {elapsed_time:.2f} seconds to create the image.")

    save_image = input("Do you want to save the image?\n")
    if save_image.lower().startswith("n") == False:
        if not os.path.exists(IMAGES_FOLDER):
            os.makedirs(IMAGES_FOLDER)
        
        image_name = input("\nWhat is the name of the image?\n")
        if image_name == "":
            image_name = created

        image_filename = f"{image_name}.png"
        image_path = image_path = os.path.join(IMAGES_FOLDER, image_filename)

        response = requests.get(image_url)
        if response.status_code == 200:
            with open(image_path, 'wb') as file:
                file.write(response.content)
            print(f"\nThe image {image_filename} was saved to {image_path}")
        else:
            print("\n Unable to save the image")
    
    # TODO
    # Include save it or not question 
