FROM golang:1.19.0-alpine

WORKDIR /app

COPY . .

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]
