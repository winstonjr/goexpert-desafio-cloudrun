FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o desafio-cloudrun ./cmd/temperature/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/desafio-cloudrun .
COPY --from=build /app/cmd/temperature/.env .
ENTRYPOINT ["./desafio-cloudrun"]
