FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor/ vendor/
COPY . .
RUN GOFLAGS=-mod=vendor go build -o /app/api ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/api .
EXPOSE 8080
CMD ["./api"]
