FROM golang:1.25-alpine AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

FROM alpine:latest AS host

WORKDIR /app

COPY --from=build /usr/local/bin/app .

EXPOSE 80
CMD ["./app"]
