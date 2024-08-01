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

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod tidy
COPY . ./
RUN task build

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/app/server /app/server

# Run the web service on container startup.
EXPOSE 3000
CMD ["/app/server"]
