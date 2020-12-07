FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

CMD ["go", "run", "main.go"]