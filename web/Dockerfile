FROM node:10.12-alpine as build-stage
WORKDIR /app
COPY . .
RUN apk update && \
yarn global add @vue/cli
