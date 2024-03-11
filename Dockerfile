FROM golang:1.22-alpine as base1

WORKDIR /app

COPY . .

RUN go build -o /main

FROM golang:1.22-alpine as base2

COPY --from=base1 /main /main

EXPOSE 8080

CMD [ "/main" ]
