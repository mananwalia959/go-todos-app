FROM golang:1.16.6-alpine3.14 AS go-build

WORKDIR /app
COPY . .

RUN go build -o main  ./cmd/web/...



FROM node:12 AS node-build
WORKDIR /app
COPY /client .

RUN npm ci

RUN npm run build

FROM golang:1.16.6-alpine3.14

WORKDIR /app

COPY --from=go-build   /app/main . 

COPY --from=node-build /app/build  client/build/

EXPOSE 8080

CMD ["main"]

