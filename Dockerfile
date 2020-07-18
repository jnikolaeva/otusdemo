FROM golang:1.14.2 as builder
LABEL maintainer="Julia N."
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/service ./cmd/service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/cli ./cmd/cli

######## Start a new stage #######
FROM alpine:3.11.5
RUN adduser -D otus
USER otus

COPY --from=builder /app/bin/ /app/bin/

WORKDIR /app/

EXPOSE 8080
CMD ["./bin/service"]