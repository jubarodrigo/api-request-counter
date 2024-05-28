FROM golang:1.22-alpine3.16

RUN apk update && apk add --no-cache gcc g++

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["/app/main"]