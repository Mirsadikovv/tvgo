FROM python:3.12.2-slim

WORKDIR /app

RUN pip install --upgrade pip

ENV PIP_TIMEOUT=100

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chmod +x start.sh

CMD ["./start.sh"]
