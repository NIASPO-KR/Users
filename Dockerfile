FROM golang:1.23.1

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o users ./cmd/users/users.go

EXPOSE 8082

CMD ["./users"]
