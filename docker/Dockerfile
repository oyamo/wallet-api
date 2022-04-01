# Initial stage: download modules
FROM golang:1.18-alpine as builder

RUN apk add build-base
RUN apk --update add git
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV config=docker

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ENV CONFIG=docker

RUN go build -o /tmp/wallet-api github.com/oyamo/wallet-api/cmd/api
RUN cp /tmp/wallet-api /usr/bin/wallet-api
RUN rm /tmp/wallet-api

# Make the ports visible
EXPOSE 5000
EXPOSE 5555
EXPOSE 7070
CMD ["wallet-api"]