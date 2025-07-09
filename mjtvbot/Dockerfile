FROM python:3.11-alpine

WORKDIR /app

RUN pip install --upgrade pip

ENV PIP_TIMEOUT=100

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chmod +x start.sh

CMD ["./start.sh"]
