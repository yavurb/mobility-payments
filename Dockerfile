FROM golang:1.23-alpine AS build

ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/

# RUN go test ./...
RUN go build -o /app/mobility-payments ./cmd/mobility-payments/main.go

FROM alpine:3.20

WORKDIR /app

RUN adduser -D mobility-payments && chown -R mobility-payments:mobility-payments /app
USER mobility-payments

# TODO: Copy configuration file
COPY --from=build --chown=mobility-payments:mobility-payments --chmod=770 /app/mobility-payments /app/

EXPOSE 8910

ENTRYPOINT ["./mobility-payments"]
