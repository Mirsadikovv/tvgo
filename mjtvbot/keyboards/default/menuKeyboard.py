from aiogram.types import ReplyKeyboardMarkup, KeyboardButton
from aiogram.types.web_app_info import WebAppInfo
menu = ReplyKeyboardMarkup(
    keyboard = [
        [KeyboardButton(text = "🆓 Бесплатные каналы", web_app=WebAppInfo(url = "https://tvgo.uz"))],
        [KeyboardButton(text = "Просмотр по подписке")],
        [KeyboardButton(text = "☑️ Мои подписки"),  KeyboardButton(text = "⚙️ Настройки")],
    ],
resize_keyboard = True)

televizor = ReplyKeyboardMarkup(
    keyboard = [
        [KeyboardButton(text = "Давайте приступим🚀")],
    ],
resize_keyboard = True)