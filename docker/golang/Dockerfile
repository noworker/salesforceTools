FROM golang:1.21.4-alpine

#PROCESS
WORKDIR /app
RUN apk update && apk add git && apk add bash
COPY ./src/go.mod ./src/go.sum ./src/.air.toml /app/
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

CMD ["air","-c",".air.toml"]