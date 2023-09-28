from golang:1.21.1 as builder

WORKDIR /app

COPY . .

RUN go build .

from ubuntu

WORKDIR /app

COPY --from=builder /app/ .

ENTRYPOINT ["./v-wallet"]