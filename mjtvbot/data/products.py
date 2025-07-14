from aiogram import types
from aiogram.types import LabeledPrice

from utils.misc.product import Product


import os


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
    photo_url=os.path.join(os.path.dirname(__file__), 'data/vip.png'),
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
    photo_url=os.path.join(os.path.dirname(__file__), 'data/sport.png'),
    photo_width=1200,
    photo_height=800,

    need_email=True,
    need_name=True,
    need_phone_number=True,
)