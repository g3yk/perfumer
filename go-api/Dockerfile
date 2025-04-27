FROM golang:1.24 AS builder

WORKDIR /compiler

COPY go.mod .

RUN go mod download && go mod verify \
  && go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

# Swagger Docs
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -o go-api ./main.go


FROM alpine:latest

WORKDIR /production

COPY --from=builder /compiler/go-api .
COPY --from=builder /compiler/docs ./docs

EXPOSE 3000

CMD ["./go-api"]
