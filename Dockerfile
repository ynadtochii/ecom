FROM golang:1.23

# Set the working directory
WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["/app/main"]

