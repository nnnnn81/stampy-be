# FROM golang:1.22-alpine
# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download && go mod verify
# EXPOSE 8080
# CMD ["go", "run", "main.go"] 

FROM golang:1.22-alpine

RUN apk update && apk add git

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY . .

RUN go install golang.org/x/tools/cmd/goimports@latest

RUN apk update && apk add git

EXPOSE 8080
CMD ["go", "run", "main.go"]