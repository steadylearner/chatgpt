# $pip install black $black -q .

from openai import OpenAI
from settings import (
  OPENAI_API_KEY,
  TEXTS_FOLDER,
  TEXT_FILE_EXT,
)
import time
import os

from termcolor import cprint

chat_gpt_model = "gpt-3.5-turbo" # gpt-4o

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

questioned_already = False
while True:
    if questioned_already == False:
        bot_message("How can I help you?")
        questioned_already = True
    else:
        bot_message("Anything else I can help?")

    # What does steadylearner mean?
    # How to make Python black package to use 2 space tab instead of 4
    print("🧑 You")
    your_message = input()

    start_time = time.time()

    chat_gpt_response = client.chat.completions.create(
        messages=[
            {
                "role": "user",
                "content": your_message,
            }
        ],
        model=chat_gpt_model,
    )

     # Stop the timer
    end_time = time.time()

    # Calculate the time difference
    elapsed_time = end_time - start_time

    text = chat_gpt_response.choices[0].message.content
    cprint(
        f"""\n🤖 Bot\n{text}""",
        # color="yellow", attrs=["bold"]
    )

    print(f"\n🤖It did take {elapsed_time:.2f} seconds to create the response.")

    save_text = input("Do you want to save it?\n")
    if save_text.lower().startswith("y") == True:
        if not os.path.exists(TEXTS_FOLDER):
            os.makedirs(TEXTS_FOLDER)
        
        text_file_name = input("\nWhat is the name of the text file?\n")
        if text_file_name == "":
            current_timestamp = int(time.time())
            text_file_name = current_timestamp

        text_filename = f"{text_file_name}.{TEXT_FILE_EXT}"
        text_file_path = os.path.join(TEXTS_FOLDER, text_filename)

        with open(text_file_path, 'w') as file:
            # file.write(f"""🤖 Bot\n{text}""")
            file.write(f"""{text}""")
            print(f"\nThe response {text_file_name} was saved to {text_file_path}")

# ChatModel = Literal[
#     "gpt-4o",
#     "gpt-4o-2024-05-13",
#     "gpt-4-turbo",
#     "gpt-4-turbo-2024-04-09",
#     "gpt-4-0125-preview",
#     "gpt-4-turbo-preview",
#     "gpt-4-1106-preview",
#     "gpt-4-vision-preview",
#     "gpt-4",
#     "gpt-4-0314",
#     "gpt-4-0613",
#     "gpt-4-32k",
#     "gpt-4-32k-0314",
#     "gpt-4-32k-0613",
#     "gpt-3.5-turbo",
#     "gpt-3.5-turbo-16k",
#     "gpt-3.5-turbo-0301",
#     "gpt-3.5-turbo-0613",
#     "gpt-3.5-turbo-1106",
#     "gpt-3.5-turbo-0125",
#     "gpt-3.5-turbo-16k-0613",
# ]
