from aiogram.dispatcher.filters import Command, Text
from aiogram.types import Message, ReplyKeyboardRemove
from keyboards.default.menuKeyboard import menu
from keyboards.default.tarifKeyboard import tarif,tarif_VIP,tarif_sport,tarif_VIPandsport
from keyboards.default.settingKeyboard import setting
from dateutil.relativedelta import relativedelta

import logging

from loader import dp,db,bot


# @dp.message_handler(text = "Давайте приступим🚀")
# async def show_menu(message: Message):
#     logging.info(message)
#     await message.delete()
#     await message.answer("Выберите одно из следующих:", reply_markup=menu)



@dp.message_handler(text='Просмотр по подписке')
async def send_link(message: Message):
    status = await db.select_user_status(telegram_id = message.from_user.id)
    for i in status:
        status = i
        break
    for j in range(1):
        if status == None:
            await message.answer("Вот наши платные тарифы",reply_markup = tarif)

        elif status == "VIP":
            status_end = await db.select_user_status_endtime(telegram_id = message.from_user.id)
            current_time = await db.select_current_time()
            for i in status_end:
                status_end = i
                # print("sddddddddddddddddd",i)
                break
            date_only = status_end.date()
            date_only = date_only + relativedelta(months=1)  # Add one month
            status_end = date_only.strftime('%d.%m.%Y')
            for j in current_time:
                for k in j:
                    current_time = k
                    # print("Dsssssssssssssssssssssssssssssssssssss",k)
                break
            date_only = current_time.date()
            current_time = date_only.strftime('%d.%m.%Y')


            if current_time>status_end:
                await db.update_user_status(status=None,telegram_id=message.from_user.id)
                await db.update_user_endtime(telegram_id=message.from_user.id)
                await message.answer("Вот наши платные тарифы",reply_markup = tarif)
                break

            await message.answer("Вот наши платные тарифы\nУ вас подписка на VIP",reply_markup = tarif_VIP)


        elif status == "sport":
            status_end = await db.select_user_status_endtime(telegram_id = message.from_user.id)
            current_time = await db.select_current_time()
            for i in status_end:
                status_end = i
                # print("sddddddddddddddddd",i)
                break
            date_only = status_end.date()
            date_only = date_only + relativedelta(months=1)  # Add one month
            status_end = date_only.strftime('%d.%m.%Y')
            for j in current_time:
                for k in j:
                    current_time = k
                    # print("Dsssssssssssssssssssssssssssssssssssss",k)
                break
            date_only = current_time.date()
            current_time = date_only.strftime('%d.%m.%Y')


            if current_time>status_end:
                await db.update_user_status(status=None,telegram_id=message.from_user.id)
                await db.update_user_endtime(telegram_id=message.from_user.id)
                await message.answer("Вот наши платные тарифы",reply_markup = tarif)
                break

            await message.answer("Вот наши платные тарифы\nУ вас подписка на спортивные каналы",reply_markup = tarif_sport)


        elif status == "VIPandsport":
            status_end = await db.select_user_status_endtime(telegram_id = message.from_user.id)
            current_time = await db.select_current_time()
            for i in status_end:
                status_end = i
                # print("sddddddddddddddddd",i)
                break
            date_only = status_end.date()
            date_only = date_only + relativedelta(months=1)  # Add one month
            status_end = date_only.strftime('%d.%m.%Y')
            for j in current_time:
                for k in j:
                    current_time = k
                    # print("Dsssssssssssssssssssssssssssssssssssss",k)
                break
            date_only = current_time.date()
            current_time = date_only.strftime('%d.%m.%Y')

            if current_time>status_end:
                await db.update_user_status(status=None,telegram_id=message.from_user.id)
                await db.update_user_endtime(telegram_id=message.from_user.id)
                await message.answer("Вот наши платные тарифы",reply_markup = tarif)
                break

            await message.answer("Вы купили все тарифы!",reply_markup = tarif_VIPandsport)        









@dp.message_handler(text='☑️ Мои подписки')
async def send_link(message: Message):
    user = await db.select_user_status(telegram_id = message.from_user.id)
    for i in user:
        if i == None:
           await message.answer("В данный момент у вас нет активных подписок\nНо вы все равно можете смотреть наши бесплатные каналы",reply_markup = menu)
        elif i == "VIP":
            ttt = await db.select_user_status_endtime(telegram_id = message.from_user.id)
            for i in ttt:
                ttt=i
                break
            date_only = ttt.date()
            ttt = date_only.strftime('%d.%m.%Y')
            time = "У вас подписка VIP\nИстекает "+str(ttt)
            await message.answer(text=time,reply_markup = menu)

        elif i == "sport":
            ttt = await db.select_user_status_endtime(telegram_id = message.from_user.id)
            for i in ttt:
                ttt=i
                break
            date_only = ttt.date()
            ttt = date_only.strftime('%d.%m.%Y')
            time = "У вас подписка Sport\nИстекает "+str(ttt)
            await message.answer(text=time,reply_markup = menu)

        elif i == "VIPandsport":
            ttt = await db.select_user_status_endtime(telegram_id = message.from_user.id)
            for i in ttt:
                ttt=i
                break
            date_only = ttt.date()
            ttt = date_only.strftime('%d.%m.%Y')
            time = "У вас подписка VIP and Sport\nИстекает "+str(ttt)
            await message.answer(text=time,reply_markup = menu)
        break


@dp.message_handler(text='⚙️ Настройки')
async def send_link(message: Message):
    await message.answer("Выберите что хотите изменить:",reply_markup = setting)

@dp.message_handler(text='Назад🔙')
async def send_link(message: Message):
    await message.answer(text = "Выберите действие:", reply_markup = menu)

