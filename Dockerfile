FROM golang:alpine
WORKDIR /app

RUN apk add --no-cache make curl gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD make all-tests-inside-container