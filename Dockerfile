FROM golang:1.22-alpine as base

WORKDIR /app

COPY . .

RUN go build -o /main

FROM golang:1.22-alpine

COPY --from=base /main /main

EXPOSE 8080

CMD [ "/main" ]
