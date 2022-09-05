FROM golang:1.19.0-alpine as builder

ENV GIN_MODE=release
ENV APP_PORT=3000

WORKDIR /app
COPY . .

RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/app ./

################################
FROM alpine:latest

RUN apk update && rm /var/cache/apk/*
RUN apk --no-cache add ca-certificates

WORKDIR /app
EXPOSE 3000

COPY --from=builder /out/app ./
CMD ["./app"]
