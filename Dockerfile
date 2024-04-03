FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV GOOS linux

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/main cmd/app/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 3180

CMD ["./main"]
