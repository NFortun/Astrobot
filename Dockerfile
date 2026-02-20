FROM golang:alpine as app-builder
WORKDIR /go/src/app
COPY . .
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o astrobot-server ./cmd/astrobot-server/main.go

FROM debian:bookworm-slim

# Get Chrome
RUN apt-get update && apt-get install -y \
    chromium \
    ca-certificates \
    fonts-liberation \
    libnss3 \
    libxss1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libgtk-3-0 \
    --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

COPY --from=app-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=app-builder /go/src/app/astrobot-server /astrobot-server
EXPOSE 8081 8082
ENTRYPOINT ["/astrobot-server"]