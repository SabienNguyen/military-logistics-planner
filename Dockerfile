# Use the official Golang image
FROM golang:1.24.4

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o military-logistics-planner ./cmd

EXPOSE 8080

CMD ["./military-logistics-planner"]
