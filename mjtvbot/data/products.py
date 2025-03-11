from aiogram import types
from aiogram.types import LabeledPrice

from utils.misc.product import Product


ds_vip = Product(
    title="VIP",
    description="Оплатите чтобы купить подписку VIP",
    currency="UZS",
    prices=[
        LabeledPrice(
            label='Подписка VIP',
            amount=2200000, 
        ),
        LabeledPrice(
            label='Скидка от админа',
            amount=-500000, 
        ),
    ],
    start_parameter="create_invoice_ds_vip",
    photo_url='https://avatars.mds.yandex.net/i?id=e4e2d6728f62814504ae458a29efaafe-4265706-images-thumbs&n=13',
    photo_width=1200,
    photo_height=800,

    need_email=True,
    need_name=True,
    need_phone_number=True,
)


ds_sport = Product(
    title="Sport",
    description="Оплатите чтобы купить подписку Sport",
    currency="UZS",
    prices=[
        LabeledPrice(
            label='Подписка на Sport',
            amount=1600000, 
        ),
        LabeledPrice(
            label='Скидка от админа',
            amount=-300000,
        ),
    ],
    start_parameter="create_invoice_ds_vip",
    photo_url='https://avatars.mds.yandex.net/i?id=aa319fae26d038194a903b6c48b0780d4054b34f-9844228-images-thumbs&n=13',
    photo_width=1200,
    photo_height=800,

    need_email=True,
    need_name=True,
    need_phone_number=True,
)