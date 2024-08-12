FROM golang:1.22-alpine3.19 as builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o bin/mdpages ./cmd/mdpages/main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/assets ./assets

RUN apk --update --no-cache add curl

EXPOSE 80
ENTRYPOINT ["./mdpages"]