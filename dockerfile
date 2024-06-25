FROM golang:1.22-alpine3.19 as builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o bin/mdpages ./main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/.env .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

RUN apk --update --no-cache add curl

EXPOSE 5000
ENTRYPOINT ["./mdpages"]