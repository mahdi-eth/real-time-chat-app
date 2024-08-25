FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o history ./cmd

EXPOSE 8081

CMD ["./history"]
