FROM golang:1.26.2-alpine3.23

ARG port

WORKDIR /app

RUN apk add build-base

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go env -w CGO_ENABLED=1 && go build forum

EXPOSE $port

CMD [ "/app/forum" ]

## an alternative to docker compose
## docker run -d --name forum -p 8080:8080 --env-file .env -v ./assets:/app/assets -v ./db:/app/db forum