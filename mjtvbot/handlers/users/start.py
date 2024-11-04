from aiogram import types
from aiogram.dispatcher.filters.builtin import CommandStart
import logging
from loader import dp,db,bot
from states.personalData import PersonalData
from aiogram.dispatcher import FSMContext
import asyncpg 
from data.config import ADMINS
from keyboards.default.tarifKeyboard import nomer
from keyboards.default.menuKeyboard import televizor
from keyboards.default.menuKeyboard import menu


@dp.message_handler(CommandStart())
async def bot_start(message: types.Message):

    try:
        user = await db.add_user(telegram_id=message.from_user.id,
                                 full_name=message.from_user.full_name,
                                 username=message.from_user.username
                                 )
    except asyncpg.exceptions.UniqueViolationError:
        user = await db.select_user(telegram_id=message.from_user.id)

    logging.info(message)
    await message.answer(f"""Привет, {message.from_user.full_name}!
Добро пожаловать в тв-бот""")
    
    # ADMINGA xabar beramiz
    count = await db.count_users()
    msg = f"{user[1]} был добавлен в базу.\nВ базе {count} пользователей."
    await bot.send_message(chat_id=ADMINS[0], text=msg)
    await message.answer(text = "Чтобы использовать бота отправьте нам номер телефона)",reply_markup = nomer)
    await PersonalData.phoneNum.set()



@dp.message_handler(content_types=[types.ContentType.CONTACT, types.ContentType.TEXT],state=PersonalData.phoneNum)
async def answer_phone(message: types.Message, state: FSMContext):

    if message.content_type == types.ContentType.CONTACT:
        contact = message.contact
        phone = contact.phone_number
        # print(phone)

        await state.update_data({"phone": phone})
        data = await state.get_data()

        phone = data.get("phone")
        await db.update_user_phone(phone = str(phone), telegram_id=message.from_user.id)
        msg = "Ваши данные:\n"
        msg += f"Имя - {message.from_user.full_name}\n"
        msg += f"Номер - {phone}"
        await message.answer(msg, reply_markup=menu)
        await state.finish()
    
    elif message.text == "Пропустить":
        await message.answer("Но имейте ввиду что без номера телефона невозможна покупка тарифов!", reply_markup=menu)
        await state.finish()

    elif message.content_type == types.ContentType.TEXT:
        await message.reply("Отправьте свой контакт пожалуйста)\nОн потребуется для дальнейшего использования бота и покупки тарифов)")
        await PersonalData.phoneNum.set()

