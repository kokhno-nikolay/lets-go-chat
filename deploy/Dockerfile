FROM golang:1.16

WORKDIR /app
COPY . .

RUN chmod +x ./scripts/wait-for-it.sh
RUN go build cmd/service/main.go
CMD ["/app/main"]