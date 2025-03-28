from aiogram import executor
from loader import dp,db
import middlewares, filters, handlers
from utils.notify_admins import on_startup_notify
from utils.set_bot_commands import set_default_commands


async def on_startup(dispatcher):
    await db.create()
    # await db.drop_users()https://tbs01-edge11.itdc.ge/ort/tracks-v1a1/mono.m3u8
    await db.create_table_users()

    await set_default_commands(dispatcher)
    await on_startup_notify(dispatcher)


if __name__ == '__main__':
    executor.start_polling(dp, on_startup=on_startup, skip_updates=False)
