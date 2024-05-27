FROM golang:1.22.3-alpine3.19

WORKDIR /usr/src/app

COPY go.* .

# this is needed to docker-compose watch reinstall everything
RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .
