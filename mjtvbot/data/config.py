from environs import Env

# environs kutubxonasidan foydalanish
env = Env()
env.read_env()

# .env fayl ichidan quyidagilarni o'qiymiz
BOT_TOKEN = env.str("BOT_TOKEN")  # Bot toekn
ADMINS = env.list("ADMINS")  # adminlar ro'yxati
IP = env.str("ip")  # Xosting ip manzili


DP_USER = env.str("DP_USER")
DP_PASS = env.str("DP_PASS")
DP_NAME = env.str("DP_NAME")
DP_HOST = env.str("DP_HOST")
PROVIDER_TOKEN = env.str("PROVIDER_TOKEN")

# import os

# # env fayl ichidan guyidagilanni a'giymiz
# BOT_TOKEN = str(os.environ.get("BOT_TOKEN")) # Bot token
# ADMINS = list(os.environ.get("ADMINS")) # adminlar rolyxati
# IP = str(os.environ.get("ip")) # Xosting ip manzili
# PROVIDER_TOKEN = str(os.environ.get ("PROVIDER_TOKEN"))