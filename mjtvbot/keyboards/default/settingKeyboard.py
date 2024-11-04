from aiogram.types import ReplyKeyboardMarkup, KeyboardButton

setting = ReplyKeyboardMarkup(
    keyboard = [
        [
            KeyboardButton(text='Профиль👤'),
            KeyboardButton(text='Язык🇺🇿🇷🇺🇬🇧'),
        ],
        [KeyboardButton(text='Назад🔙')]
    ],
    resize_keyboard=True
)

admin_panel = ReplyKeyboardMarkup(
    keyboard = [
        [
            KeyboardButton(text='Подписка 🪓'),
            KeyboardButton(text='Tanish bilish😎'),
        ],
        [KeyboardButton(text='Реклама📣'),KeyboardButton(text='База пользователей📊')],
        [KeyboardButton(text='Назад🔙')]
    ],
    resize_keyboard=True
)

reklama = ReplyKeyboardMarkup(
    keyboard = [
        [
            # KeyboardButton(text='Создать рекламу'),
            KeyboardButton(text='Рассылка📢'),
        ],
        [KeyboardButton(text='Назад')]
    ],
    resize_keyboard=True
)

nazad = ReplyKeyboardMarkup(
    keyboard = [
        [KeyboardButton(text='Назад')]
    ],
    resize_keyboard=True
)