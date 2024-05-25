<!-- # How to use ChatGPT with Python -->

[Python website]: https://www.python.org
[ChatGPT]: https://chatgpt.com
[openai-python]: https://github.com/openai/openai-python
[openai-api-keys]: https://platform.openai.com/account/api-keys
[openai-api-limits]: https://platform.openai.com/settings/organization/limits 
[openai-api-billing]: https://platform.openai.com/settings/organization/billing
[openai-api-usage]: https://platform.openai.com/usage 

[Steadylearner GitHub]: https://github.com/steadylearner
[Steadylearner ChatGPT repository]: https://github.com/steadylearner/chatgpt

[Hire me]: https://www.onlycoiners.com/user/steadylearner/product/i-will-help-you-to-build-a-full-stack-app-bot-or-smart-contr

[ChatGPT YouTube]: https://www.youtube.com/watch?v=JY7H286Jubg

In this post, you will learn how to use ChatGPT with Python using the official [openai-python] package. 

[You can find the code used here in this repository.][Steadylearner ChatGPT repository] It also has examples in other programming languages and blog posts for them.

[If you haven't used ChatGPT for images, you might find this useful.][ChatGPT YouTube]

I plan to include more examples with other programming languages, so feel free to like the repository if you find it useful and want to see more examples.

[You can also hire me if you need a full stack developer for your projects.][Hire me]

## Prerequisites

You can skip these and go to **Code** if you have done similar things already.

1. Install Python
2. Set up virtual envrionment
3. Get OpenAI API key
4. Install Python packages

### 1. Install Python

If you don't already have Python installed, visit the [Python website] and follow the guide to download it. To check if Python is installed, use your console and type this.

```console
$python --version
```

Depending on OS you use, you might see an instruction to install it.

### 2. Set up virtual environment

Create a Python virtual environment to manage your dependencies. Here, we will name the environment chatgpt, but you can choose any name

```console
$python3 -m venv chatgpt
```

Then, you will see these are created in your new chatgpt folder.

```console
bin		include		lib		pyvenv.cfg
```

Use `$cd chatgpt` and `$source bin/activate` to activate your Python development environment.

### 3. Get OpenAI API key

Sign up on the OpenAI platform and get your API key from the [API keys page][openai-api-keys]. Store it securely as you will need it to access the OpenAI API later.

You can find more information about OpenAI API usage, billing, and limits at the following links.

- [API Limits][openai-api-limits]
- [API Billing][openai-api-billing]
- [API Usage][openai-api-usage]

### 4. Install Python packages

You can install pacakges required for this post with this.

```console
$pip install openai requests termcolor
```

If you want to use requirements.txt in the repository instead, you can use this.

```console
$pip install -r requirements.txt
```

To create your own requirements.txt file, you can use this.

```console
$pip freeze > requirements.txt
```

## Code

### Configuration

First, create a .env file in your chatgpt folder with the following text.

```console
OPENAI_API_KEY=<YOURS>
```

Then, create a settings.py file to manage your settings.

```py
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
QUALITY = "hd"
```

Finally, we will see two examples that you might find useful and make some variations later.

### A CLI chatbot

First, we will see a simple code snippet for a CLI chatbot.

You can save the text below as **text.py** or with another name.

```py
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
    api_key=OPENAI_API_KEY,
)

bot_message = lambda text: cprint(
    f"""
🤖 Bot 
{text}
""",
)

questioned_already = False
while True: # 1.
    if questioned_already == False:
        bot_message("How can I help you?")
        questioned_already = True
    else:
        bot_message("Anything else I can help?")

    print("🧑 You")
    your_message = input()

    # 2.
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

    # 2.
    end_time = time.time()
    elapsed_time = end_time - start_time

    text = chat_gpt_response.choices[0].message.content
    cprint(
        f"""\n🤖 Bot\n{text}""",
    )

    print(f"\n🤖It did take {elapsed_time:.2f} seconds to create the response.")

    save_text = input("Do you want to save it?\n")
    if save_text.lower().startswith("y") == True:
        # 3.
        if not os.path.exists(TEXTS_FOLDER):
            os.makedirs(TEXTS_FOLDER)
        
        text_file_name = input("\nWhat is the name of the text file?\n")
        # 4.
        if text_file_name == "":
            current_timestamp = int(time.time())
            text_file_name = current_timestamp

        text_filename = f"{text_file_name}.{TEXT_FILE_EXT}"
        text_file_path = os.path.join(TEXTS_FOLDER, text_filename)

        with open(text_file_path, 'w') as file:
            # file.write(f"""🤖 Bot\n{text}""")
            file.write(f"""{text}""")
            print(f"\nThe response {text_file_name} was saved to {text_file_path}")
```

Then, use `$python text.py` and you will see some messages like this.

You can use your own question instead of "What does steadylearner mean?"

```console
🤖 Bot 
How can I help you?

🧑 You
What does steadylearner mean?

🤖 Bot
Steadylearner likely means someone who is constantly learning and seeking knowledge in a consistent and steady manner. This person is dedicated to their personal growth and development through continuous learning and improvement.

Do you want to save it?
n

🤖 Bot 
Anything else I can help?
```

You can also optionally save a response from the bot to **texts** folder you used at settings.py file.

So how does it work? There are a few important parts but we will see only important parts.

1. We want to ask questions to the bot until we stop the bot with `Ctrl + C` or others that results in SIGINT signal.
2. The time for the ChatGPT API to respond varies depending on your question and it might take very long so we will have some feedbacks by measuring how much time it takes.
3. We make a folder to save the bot response in case it didn't exist before.
4. You might not want to type file name everytime so we use timestamp for a file name by default.

So that was the simple code snippet needed to respond to your questions.

The **gpt-3.5-turbo** option was used here to save resources but you can also use **gpt-4o** option.

Not so interesting because you can use more features at [ChatGPT] website. Therefore, we will make something a little bit more useful variation.

### A CLI image creator

The code snippet used here is almost the same to the previous one but we will use **dall-e-3** model to make images. You can use it for profile image, blog cover, logo etc.

```py
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
    api_key=OPENAI_API_KEY,
)

bot_message = lambda text: cprint(
    f"""
🤖 Bot 
{text}
""",
)

while True:
    bot_message("Can you describe the image you want to create with details?")

    print("🧑 You")
    your_message = input()

    start_time = time.time()

    chat_gpt_response = client.images.generate(
        model=chat_gpt_model,
        prompt=your_message,
        size=IMAGE_SIZE,
        quality=QUALITY,
        n=1,
    )

    end_time = time.time()

    elapsed_time = end_time - start_time

    created = chat_gpt_response.created
    revised_prompt = chat_gpt_response.data[0].revised_prompt
    image_url = chat_gpt_response.data[0].url

    cprint(
        f"""\n🤖 Bot\nHere is the link to the image.\n\n{image_url}\n""",
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
```

You can test it with `$python image.py` command. Then, describe the image you want to create with details.

You will see similar messages you read in the previous process but you will see an image created with url and how much time it did take.

You can use another message instead of **Can you make a cover image for blog post with title "How to use chatgpt with Python"** used here.

```console
🤖 Bot 
Can you describe the image you want to create with details?

🧑 You
Can you make a cover image for blog post with title "How to use chatgpt with Python"

🤖 Bot
Here is the link to the image.

https://oaidalleapiprodscus.blob.core.windows.net/private

Do you want to save the image?
y

What is the name of the image?
cover

The image cover.png was saved to images/cover.png
```

The profile and the cover image on this blog was created with it. 

You can also make your own images.

## Conclusion

You have learned how to create simple CLIs for text and image using Python and [openai-python]. 

You can extend these examples to create more complex applications. 

If you thought this post helpful, please share it and like the [ChatGPT repository][Steadylearner ChatGPT repository] and blog post.

[You can also hire me if you need a full stack developer for your projects.][Hire me]