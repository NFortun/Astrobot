FROM golang:alpine as app-builder
WORKDIR /go/src/app
COPY . .
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o astrobot-server ./cmd/astrobot-server/main.go

FROM scratch
COPY --from=app-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=app-builder /go/src/app/astrobot-server /astrobot-server
COPY --from=app-builder /go/src/app/config.json /config.json
EXPOSE 3000
ENTRYPOINT ["/astrobot-server"]