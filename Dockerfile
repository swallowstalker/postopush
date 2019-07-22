FROM python:3.6.9-stretch

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN pip install -r requirements.txt
ENTRYPOINT python push.py