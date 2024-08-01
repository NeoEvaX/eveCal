FROM golang:1.22.5-bullseye AS builder

WORKDIR /app

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest ./
RUN go install github.com/go-task/task/v3/cmd/task@latest ./

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN task build

FROM scratch
WORKDIR /app
COPY --from=builder /build .
EXPOSE 3000
CMD [ "/build/server" ]
