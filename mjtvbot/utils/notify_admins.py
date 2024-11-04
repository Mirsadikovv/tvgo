import logging

from aiogram import Dispatcher

from data.config import ADMINS


async def on_startup_notify(dp: Dispatcher):
    photo = "https://avatars.mds.yandex.net/i?id=f97c1c9da4b2706eec172e9e009e9ee3-5870104-images-thumbs&n=13"
    for admin in ADMINS:
        try:
            await dp.bot.send_photo(admin,photo,caption="Бот начал работу  /start\n\nКоманды для админов:\n/admin - админ панель")

        except Exception as err:
            logging.exception(err)
