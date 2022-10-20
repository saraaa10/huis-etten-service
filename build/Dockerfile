FROM golang:1.19.2-alpine3.16

RUN apk add git
WORKDIR /src/app

COPY go.mod go.sum ./
RUN go.mod download && go mod verify
COPY . .

RUN go build -o /src/app
CMD ["/src/app"]