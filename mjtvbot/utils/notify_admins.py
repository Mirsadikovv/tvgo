import logging

from aiogram import Dispatcher

from data.config import ADMINS


async def on_startup_notify(dp: Dispatcher):
    photo = "https://repository-images.githubusercontent.com/735638594/8a5e89ee-83e4-40c9-b735-5cedfe4901f7"
    for admin in ADMINS:
        try:
            await dp.bot.send_photo(admin,photo,caption="Бот начал работу  /start\n\nКоманды для админов:\n/admin - админ панель")

        except Exception as err:
            logging.exception(err)
