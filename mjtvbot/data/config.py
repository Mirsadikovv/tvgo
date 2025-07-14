from environs import Env

env = Env()
env.read_env()

BOT_TOKEN = env.str("BOT_TOKEN")
ADMINS = env.list("ADMINS")
IP = env.str("ip")


DP_USER = env.str("DP_USER")
DP_PASS = env.str("DP_PASS")
DP_NAME = env.str("DP_NAME")
DP_HOST = env.str("DP_HOST")
PROVIDER_TOKEN = env.str("PROVIDER_TOKEN")

