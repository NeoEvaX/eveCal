FROM golang:1.22.5-bullseye AS builder

WORKDIR /app

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_current.x | bash - && \
	apt-get install -y nodejs \
	build-essential && \
	node --version && \ 
	npm --version

# Install Templ
RUN go install github.com/a-h/templ/cmd/templ@latest 

# Install Task
RUN go install github.com/go-task/task/v3/cmd/task@latest 

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN task build

FROM scratch
WORKDIR /build
COPY --from=builder /build .
EXPOSE 3000
CMD [ "/build/server" ]
