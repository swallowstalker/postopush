import telegram
import os


def main():
    token = os.getenv("TOKEN", None)
    message = os.getenv("MESSAGE", "No message, please set MESSAGE env")
    chat_id = os.getenv("CHAT_ID", None)

    bot = telegram.Bot(token=token)
    bot.send_message(chat_id=chat_id, text=message, parse_mode=telegram.ParseMode.HTML)


if __name__ == "__main__":
    main()