FROM golang:1.22-alpine AS build
WORKDIR /app
# cache buster - her buildta değişir
ARG CACHE_BUST=1
RUN echo "cache-bust-$CACHE_BUST" && mkdir -p /tmp/clear-cache
# go.sum ve go.mod dosyalarını kopyala
COPY go.mod go.sum ./
COPY vendor/ vendor/
# tüm kaynak kodunu kopyala
COPY . .
RUN GOFLAGS=-mod=vendor go build -o /app/api ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/api .
EXPOSE 8080
CMD ["./api"]
