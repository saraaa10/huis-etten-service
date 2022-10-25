FROM golang:1.19.2-alpine3.16

WORKDIR /app

COPY go.mod go.sum ./
RUN go.mod download && go mod verify
COPY . .

RUN go build -o main .
CMD ["/app"]