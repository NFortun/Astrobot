FROM golang:alpine as app-builder
WORKDIR /go/src/app
COPY . .
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o astrobot-server ./cmd/astrobot-server/main.go

FROM ubuntu:latest

# Get Chrome
RUN apt-get update && apt-get install -y wget gnupg2 snapd
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list
RUN apt-get update
RUN apt-get install chromium
    
COPY --from=app-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=app-builder /go/src/app/astrobot-server /astrobot-server
EXPOSE 8080 8081
ENTRYPOINT ["/astrobot-server"]