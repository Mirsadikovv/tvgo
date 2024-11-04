from aiogram import types
from keyboards.default.menuKeyboard import menu
from loader import dp
import logging

# Echo bot
# @dp.message_handler(text = "/nmabu")
# async def photo_id(message: types.Message):
#     await dp.bot.send_photo(chat_id=message.from_user.id,photo = "AgACAgQAAxkDAAIEeGU6b7vgG6M5hF7GVXRFpaX1yH6dAAI0sTEbkotlUIK0Nwu42taAAQADAgADcwADMAQ")



@dp.message_handler(state=None)
async def bot_echo(message: types.Message):
    logging.info(message)
    await message.answer("Выберите действие",reply_markup=menu)

@dp.message_handler(state=None, content_types=types.ContentType.PHOTO)
async def photo_id(message: types.Message):
    await message.answer(message.photo[-1].file_id)


