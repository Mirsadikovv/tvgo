from aiogram import types
from aiogram.dispatcher.filters.builtin import CommandHelp

from loader import dp


@dp.message_handler(CommandHelp())
async def bot_help(message: types.Message):
    text = ("Запросы: ",
            "/start - Запустить бота",
            "/help - Помощь",
            "Обратная связь - @hajimen1_bot")
    
    await message.answer("\n".join(text))
