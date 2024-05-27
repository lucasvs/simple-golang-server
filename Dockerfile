FROM golang:1.22.3-alpine3.19

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential

WORKDIR /usr/src/app

COPY go.* .

# this is needed to docker-compose watch reinstall everything
RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .
