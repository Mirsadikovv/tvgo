import asyncio
from aiogram.types import Message
from aiogram import types
from aiogram.types import InputFile
from data.config import ADMINS
from loader import dp, db, bot
from keyboards.default.settingKeyboard import admin_panel, reklama,nazad
from states.personalData import tg_id
from aiogram.dispatcher import FSMContext
from aiogram.utils.exceptions import ChatNotFound

# @dp.message_handler(text="/stat", user_id=ADMINS)
# async def send_statistic_to_admins(message: types.Message):
#     file1 = open("baza.txt","w")
#     users = await db.select_all_users()
#     for i in users:
#         file1.write(f"{i}\n\n")
#     count = await db.count_users()
#     with open('baza.txt', 'rb') as file1:
#         await bot.send_document(message.from_user.id, file1,caption =f"В данный момент в нашей базе {count} пользователя")
        
   

@dp.message_handler(text="/admin", user_id=ADMINS)
async def to_admin(message: types.Message):
    await bot.send_message(message.from_user.id,text="Здарова хозяин)",reply_markup=admin_panel)


@dp.message_handler(text="База пользователей📊", user_id=ADMINS)
async def to_admin(message: types.Message):
    file1 = open("baza.txt","w")
    users = await db.select_all_users()
    for i in users:
        file1.write(f"{i}\n\n")
    count = await db.count_users()
    with open('baza.txt', 'rb') as file1:
   
     await bot.send_document(message.from_user.id, file1,caption =f"В данный момент в нашей базе {count} пользователя")

@dp.message_handler(text='Tanish bilish😎', user_id=ADMINS)
async def send_link(message: Message, state: FSMContext):
    await message.answer("Отправьте telegram_id пользователя",reply_markup = nazad)
    await tg_id.state1.set()



@dp.message_handler(state=tg_id.state1)
async def send_link(message: Message, state: FSMContext):
    
    users = await db.select_all_users()
    temp = False
    for i in users:
        if message.text == "Назад":
            await bot.send_message(message.from_user.id,text="Че передумал?🤣",reply_markup=admin_panel)
            await state.finish()
            break
        elif message.text == str(i[3]):
            temp = True
            break

        
    if temp:
        await db.update_user_status(telegram_id = int(i[3]),status = 'VIPandsport')
        await db.update_user_status_date(telegram_id = int(i[3]))
        await bot.send_message(message.from_user.id,"Успешно подарили подписку",reply_markup=admin_panel)
        
        await state.finish()
    elif temp == False and  message.text != "Назад":
        await bot.send_message(message.from_user.id,"Такого пользователя нет в базе",reply_markup=admin_panel)
        await state.finish()


            


@dp.message_handler(text='Подписка 🪓', user_id=ADMINS)
async def send_link(message: Message, state: FSMContext):
    await message.answer("Отправьте  telegram_id пользователя",reply_markup = nazad)
    await tg_id.state2.set()

@dp.message_handler(state=tg_id.state2)
async def send_link(message: Message, state: FSMContext):
    users = await db.select_all_users()
    temp = False
    for i in users:
        if message.text == "Назад":
            await bot.send_message(message.from_user.id,text="Че передумал?😏",reply_markup=admin_panel)
            await state.finish()
            break
        elif message.text == str(i[3]):
            temp = True
            break
        
        
    if temp:
        await db.update_user_status(telegram_id = int(i[3]),status = None)
        await db.update_user_endtime(telegram_id = int(i[3]))
        await bot.send_message(message.from_user.id,"Успешно удалили подписку",reply_markup=admin_panel)
        await state.finish()
    elif temp == False and  message.text != "Назад":
        await bot.send_message(message.from_user.id,"Такого пользователя нет в базе",reply_markup=admin_panel)
        await state.finish()

            




@dp.message_handler(text='Реклама📣', user_id=ADMINS)
async def send_link(message: Message):
    await message.answer("Выберите действие:",reply_markup = reklama)

@dp.message_handler(text='Назад', user_id=ADMINS)
async def send_link(message: Message):
    await message.answer("Нма гап энди админ:",reply_markup = admin_panel)



# @dp.message_handler(text='Создать рекламу', )
# async def send_link(message: Message):
#     await message.answer("Отправьте фото:",reply_markup = admin_panel)



@dp.message_handler(text="Рассылка📢", user_id=ADMINS)
async def send_ad_to_all(message: types.Message):
    users = await db.select_all_users()
    photo_file = "AgACAgIAAxkBAAIKRWVBL9kiMOC-qL5GVel5ZhsTjVwBAALG0jEbRMqgSDN8vC1W8TqgAQADAgADeAADMwQ"
    # photo_file = "AgACAgIAAxkBAAIIsWU-dNV3mjY6QFi1hGo_ZBI7nNydAAJlzjEbVCjwSaQlTkeNrx7nAQADAgADeQADMAQ"
    
    # print(photo_file)
    file2 = open("blocked_baza.txt","w")
    # file2.write(f"hf,jnftn")
    block_count = 0
    for user in users:
        try:
            user_id = user[3]
            print(user_id,"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu")
            await bot.send_photo(photo = photo_file, chat_id=user_id, caption="<a href = 'https://t.me/hajimeN1'>Подпишись!</a>")
            await asyncio.sleep(0.05)
        except ChatNotFound:
            block_count+=1
            file2.write(f"{user}\n\n")
    with open('blocked_baza.txt', 'rb') as file2:
        await bot.send_document(message.from_user.id, file2,caption =f"В данный момент бота заблокировали {block_count} пользователей")
            