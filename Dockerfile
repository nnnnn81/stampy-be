FROM golang:1.22-alpine
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]

# FROM golang:1.21.4-alpine
# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# CMD ["go", "run", "main.go"] 
