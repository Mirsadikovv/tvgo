FROM python:3.11-alpine

# Установите необходимые системные пакеты
# RUN apk add --no-cache gcc musl-dev

# Установите pip
RUN pip install --upgrade pip

# Увеличьте тайм-аут
ENV PIP_TIMEOUT=100

# Копируйте файлы requirements.txt
COPY requirements.txt .

# Установите зависимости
RUN pip install --no-cache-dir -r requirements.txt


# Копируйте ваш проект
COPY . .

RUN chmod +x start.sh

# Запустите ваше приложение
CMD ["./start.sh"]
