FROM golang:1.23-alpine AS build

ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/

# RUN go test ./...
RUN go build -o /app/mobility ./cmd/mobility/main.go

FROM alpine:3.20

WORKDIR /app

RUN adduser -D mobility && chown -R mobility:mobility /app
USER mobility

# TODO: Copy configuration file
COPY --from=build --chown=mobility:mobility --chmod=770 /app/mobility /app/

EXPOSE 8910

ENTRYPOINT ["./mobility"]
