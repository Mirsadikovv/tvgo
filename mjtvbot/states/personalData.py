from aiogram.dispatcher.filters.state import StatesGroup, State


# Shaxsiy ma'lumotlarni yig'sih uchun PersonalData holatdan yaratamiz
class PersonalData(StatesGroup):
    stst = State()
    fullName = State() # ism
    email = State() # email
    phoneNum = State() # Tel raqami
    menuState = State()

class tg_id(StatesGroup):
    state1 = State()
    state2 = State() # ism

class invoys(StatesGroup):
    state11 = State()
    state22 = State() # ism