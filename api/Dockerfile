FROM golang:1.16-alpine

ARG DbURL

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN echo 'DbURL: '$DbURL > config.yaml

RUN go mod download

EXPOSE 3333

CMD ["go", "run", "main.go"]
