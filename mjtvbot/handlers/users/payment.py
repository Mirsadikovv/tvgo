from aiogram import types
from aiogram.dispatcher.filters import Command
from aiogram.types import CallbackQuery, Message
from data.config import ADMINS

from loader import dp,db,bot
from data.products import ds_vip, ds_sport
from keyboards.inline.product_keys import build_keyboard
from keyboards.default.menuKeyboard import menu
from states.personalData import invoys
from aiogram.dispatcher import FSMContext


# @dp.message_handler(Command("VIP"))
# async def show_invoices(message: types.Message):
#     caption = "<b>Подписка VIP</b> .\n\n"
#     caption += "Цена: <b>25 000 сум</b>\n"
#     caption += "Скидка от админа: <b>-6 000 сум</b>"
#     await message.answer_photo(photo="https://avatars.mds.yandex.net/i?id=e4e2d6728f62814504ae458a29efaafe-4265706-images-thumbs&n=13",
#                          caption=caption, reply_markup=build_keyboard("VIP"))

# @dp.message_handler(Command("sport"))
# async def show_invoices(message: types.Message):
#     caption = "<b>Подписка Sport</b> .\n\n"
#     caption += "Цена: <b>20 000 сум</b>\n"
#     caption += "Скидка от админа: <b>-5 000 сум</b>"
#     await message.answer_photo(photo="https://avatars.mds.yandex.net/i?id=aa319fae26d038194a903b6c48b0780d4054b34f-9844228-images-thumbs&n=13",
#                          caption=caption, reply_markup=build_keyboard("sport"))



@dp.callback_query_handler(text="product:VIP")
async def VIP_plus_invoice(call: CallbackQuery,state: FSMContext):
    invo1 = await bot.send_invoice(chat_id=call.from_user.id,
                           **ds_vip.generate_invoice(),
                           payload="payload:VIP")
    await state.update_data({"invo1": invo1.message_id})
    await call.answer()

@dp.callback_query_handler(text="product:sport")
async def sport_invoice(call: CallbackQuery, state: FSMContext):
    invo2 = await bot.send_invoice(chat_id=call.from_user.id,
                           **ds_sport.generate_invoice(),
                           payload="payload:sport")
    await state.update_data({"invo2": invo2.message_id})   
    await call.answer()

# await state.update_data({"phone": phone})
#         data = await state.get_data()

@dp.pre_checkout_query_handler()
async def process_pre_checkout_query(pre_checkout_query: types.PreCheckoutQuery, state: FSMContext):
    await bot.answer_pre_checkout_query(pre_checkout_query_id=pre_checkout_query.id,
                                        ok=True)
    await bot.send_message(chat_id=pre_checkout_query.from_user.id,
                           text="Спасибо за покупку!",reply_markup=menu)
    data = await state.get_data()
    print(data)
    # print(data[0])
    # print(data["invo2"])
    if pre_checkout_query.invoice_payload == 'payload:sport':
        await dp.bot.delete_message(chat_id=pre_checkout_query.from_user.id,message_id=data["invo2"])
    elif pre_checkout_query.invoice_payload == 'payload:VIP':
        await dp.bot.delete_message(chat_id=pre_checkout_query.from_user.id,message_id=data["invo1"])
    await bot.send_message(chat_id=ADMINS[0],
                           text=f"Куплено: {pre_checkout_query.invoice_payload}\n"
                                f"ID: {pre_checkout_query.id}\n"
                                f"Пользователь: {pre_checkout_query.from_user.first_name}\n"                                
                                f"Покупатель: {pre_checkout_query.order_info.name}, тел: {pre_checkout_query.order_info.phone_number}\n\n"
                                f"Telegram id: {pre_checkout_query.from_user.id}\n")
    await db.update_user_phone(telegram_id=pre_checkout_query.from_user.id,phone=pre_checkout_query.order_info.phone_number)
    status = await db.select_user_status(telegram_id = pre_checkout_query.from_user.id)

    for i in status:
        status = i
        break
    # print("Dsfffffffffffffff",status,pre_checkout_query.invoice_payload)

    if pre_checkout_query.invoice_payload == 'payload:sport' and status == "VIP":
        await db.update_user_status(status="VIPandsport",telegram_id=pre_checkout_query.from_user.id)
        await db.update_user_status_date(telegram_id = pre_checkout_query.from_user.id)

    elif pre_checkout_query.invoice_payload == 'payload:VIP' and status == "sport":
        await db.update_user_status(status="VIPandsport",telegram_id=pre_checkout_query.from_user.id)
        await db.update_user_status_date(telegram_id = pre_checkout_query.from_user.id)
    
    elif pre_checkout_query.invoice_payload == 'payload:sport':
        await db.update_user_status(status="sport",telegram_id=pre_checkout_query.from_user.id)
        await db.update_user_status_date(telegram_id = pre_checkout_query.from_user.id)

    elif pre_checkout_query.invoice_payload == 'payload:VIP':
        await db.update_user_status(status="VIP",telegram_id=pre_checkout_query.from_user.id)
        await db.update_user_status_date(telegram_id = pre_checkout_query.from_user.id)

    
    
    