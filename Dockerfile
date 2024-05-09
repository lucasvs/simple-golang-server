FROM golang:1.21.0-bookworm

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential

WORKDIR /usr/src/app

COPY go.* .

# this is needed to docker-compose watch reinstall everything
RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .