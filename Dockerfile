# syntax=docker/dockerfile:1

FROM golang:1.25.5

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build /app/cmd/server/server.go

EXPOSE 1323

CMD ["/app/server"]
