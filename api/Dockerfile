FROM node:lts-alpine

RUN mkdir -p /app/node_modules

WORKDIR /app

COPY package.json ./

RUN apk add --no-cache git

COPY . /app

RUN chown -R node:node /app

EXPOSE 3333

ENTRYPOINT ["node","ace","serve","--watch"]
