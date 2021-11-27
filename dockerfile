FROM golang

WORKDIR /app

COPY . .

RUN go build ./cmd/term-backend

CMD ["./term-backend"]