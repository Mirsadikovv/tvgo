from aiogram.dispatcher.filters import Command, Text
from aiogram.types import Message, ReplyKeyboardRemove
from keyboards.inline.product_keys import build_keyboard
from keyboards.default.menuKeyboard import menu
import logging
from loader import dp


# @dp.message_handler(text='VIP')
# async def send_link(message: Message):
#     photo = "AgACAgIAAxkBAAIBoGgbnuXe2GCs1I6_LRCbToNfHnT_AAIu7zEb1xrhSAMVU-4wX99eAQADAgADeAADNgQ"
#     caption =  "В тариф VIP входит\n\n"
    
#     caption += "Цена: <b>22 000 сум</b>\n"
#     caption += "Скидка от админа: <b>-5 000 сум</b>\n"
#     caption += "Итого к оплате: 17 000\n\n"
#     caption += "Для покупки нажмите👇" 
#     await message.answer_photo(photo = photo,caption = caption,reply_markup=build_keyboard("VIP"))

@dp.message_handler(text='VIP')
async def send_link(message: Message):
    photo_path = "vip.png"  # если файл лежит рядом с этим скриптом
    caption = (
        "В тариф VIP входит\n\n"
        "Цена: <b>22 000 сум</b>\n"
        "Скидка от админа: <b>-5 000 сум</b>\n"
        "Итого к оплате: 17 000\n\n"
        "Для покупки нажмите"
    )

    try:
        with open(photo_path, "rb") as photo:
            await message.answer_photo(
                photo=photo,
                caption=caption,
                reply_markup=build_keyboard("VIP")
            )
    except FileNotFoundError:
        await message.answer(caption + "\n\nОшибка отправки фото")

@dp.message_handler(text='Sport')
async def send_link(message: Message):
    photo = "AgACAgIAAxkBAAIBomgbnxMqDrIBd1WybWWKccDWS64TAAKe8zEbv8bhSBrUvO625YPgAQADAgADeAADNgQ"
    caption =  "В тариф Sport входит\n\n"
    
    caption += "Цена: <b>16 000 сум</b>\n"
    caption += "Скидка от админа: <b>-3 000 сум</b>\n"
    caption += "Итого к оплате: 13 000\n\n"
    caption += "Для покупки нажмите👇" 
    await message.answer_photo(photo = photo,caption = caption,reply_markup=build_keyboard("sport"))


@dp.message_handler(text='Назад🔙')
async def send_link(message: Message):
    logging.info(message)
    await message.answer("Выберите одно из следующих:",reply_markup = menu)
