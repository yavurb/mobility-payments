FROM golang:1.23-alpine AS build

LABEL org.opencontainers.image.source=https://github.com/yavurb/mobility-payments
LABEL org.opencontainers.image.licenses=MIT

ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/
COPY config config/

RUN go build -o /app/mobility-payments ./cmd/mobility/main.go

FROM alpine:3.20

WORKDIR /app

ARG PKL_VERSION=0.27.0

RUN apk add --no-cache curl \
  && curl -L -o /usr/local/bin/pkl "https://github.com/apple/pkl/releases/download/${PKL_VERSION}/pkl-alpine-linux-amd64" \
  && chmod +x /usr/local/bin/pkl

RUN adduser -D mobility-payments && chown -R mobility-payments:mobility-payments /app
USER mobility-payments

COPY --chown=mobility-payments:mobility-payments --chmod=440 config/ConfigSchema.pkl /app/config/
COPY --chown=mobility-payments:mobility-payments --chmod=440 config/*-config.pkl /app/config/
COPY --from=build --chown=mobility-payments:mobility-payments --chmod=770 /app/mobility-payments /app/

EXPOSE 8910

ENTRYPOINT ["./mobility-payments"]
