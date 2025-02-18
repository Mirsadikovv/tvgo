import asyncio
import nest_asyncio
from postgresql import Database
# from utils.db_api.postgresql import Database

nest_asyncio.apply() 

async def test():
    db = Database()
    await db.create()
    await db.alter_table()
    print("Users jadvalini yaratamiz...")
    # await db.drop_users()
    await db.create_table_users()
    print("Yaratildi")

    # await db.drop_users()
    print("ooooooooooooooooooooooooooooooooo")
    # print("Foydalanuvchilarni qo'shamiz")

    # await db.add_user("anvar", "sariqdev", 123456789)
    # await db.add_user("olim", "olim223", 12341123)
    # await db.add_user("1", "1", 131231)
    # await db.add_user("1", "1", 23324234)
    # await db.add_user("John", "JohnDoe", 4388229)2
    # print("Qo'shildi")

    # users = await db.select_all_users()
    # print(f"Barcha foydalanuvchilar: {users}")

    # user = await db.select_user(id=5)
    # print(f"Foydalanuvchi: {user}")

    async def alter_table(self):
        sql = "ALTER TABLE users ADD COLUMN subscription_date TIMESTAMP"
        return await self.execute(sql)


async def main():
    await test()

asyncio.run(main())

# asyncio.run(test())
# asyncio.run(test())