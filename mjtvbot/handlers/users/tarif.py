from aiogram.dispatcher.filters import Command, Text
from aiogram.types import Message, ReplyKeyboardRemove
from keyboards.inline.product_keys import build_keyboard
from keyboards.default.menuKeyboard import menu
import logging
from loader import dp


@dp.message_handler(text='VIP')
async def send_link(message: Message):
    photo = "AgACAgIAAxkBAAIZAAFnAAEhD-4LIso4_l34BnwbfkWs5ZYAAinjMRvkngABSHZ08di-3AJYAQADAgADeQADNgQ"
    caption =  "В тариф VIP входит\n\n"
    
    caption += "Цена: <b>22 000 сум</b>\n"
    caption += "Скидка от админа: <b>-5 000 сум</b>\n"
    caption += "Итого к оплате: 17 000\n\n"
    caption += "Для покупки нажмите👇" 
    await message.answer_photo(photo = photo,caption = caption,reply_markup=build_keyboard("VIP"))

@dp.message_handler(text='Sport')
async def send_link(message: Message):
    photo = "AgACAgIAAxkBAAIZAAFnAAEhD-4LIso4_l34BnwbfkWs5ZYAAinjMRvkngABSHZ08di-3AJYAQADAgADeQADNgQ"
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