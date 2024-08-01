FROM golang:1.22.5-bullseye AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /server cmd/eveCal/main.go

COPY --from=builder /server /server

ENV PORT=8080
EXPOSE $PORT

ENTRYPOINT ["/server"]
