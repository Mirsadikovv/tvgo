from aiogram import types

from environs import Env

env = Env()
env.read_env()

admins = env.str("ADMINS")

async def set_default_commands(dp):
    current_user = await dp.bot.get_me()
    if current_user in admins:
        await dp.bot.set_my_commands(
            [
                types.BotCommand("start", "Начать работу"),
                types.BotCommand("help", "Помощь"),
                types.BotCommand("admin", "Админ панель"),
            ]
        )
    else:
        await dp.bot.set_my_commands(
            [
                types.BotCommand("start", "Начать работу"),
                types.BotCommand("help", "Помощь"),
            ]
        )

