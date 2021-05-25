FROM golang:alpine AS builder

ARG APP_NAME

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY api ./api
RUN go build -o app ./cmd/${APP_NAME}

FROM node:alpine AS builder-f2e

WORKDIR /src

COPY web/package.json ./
RUN yarn install

COPY web/public ./public
COPY web/src ./src
RUN yarn build

FROM alpine:3 AS final

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/app ./
COPY --from=builder-f2e /src/build ./web/build

ENTRYPOINT ./app