FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./callback-server .


FROM alpine:3

COPY --from=builder /app/callback-server ./
RUN chmod +x ./callback-server

EXPOSE 8080
ENTRYPOINT ["./callback-server"]