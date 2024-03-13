FROM golang:1.22-alpine
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]

# FROM golang:1.21.4-alpine
# WORKDIR /app

# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o binary

# EXPOSE 8080

# CMD ["./binary"]
