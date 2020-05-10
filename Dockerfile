FROM golang:1.14.2 as builder
LABEL maintainer="Julia N."
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/otusdemo ./cmd

######## Start a new stage #######
FROM alpine:3.11.5
RUN adduser -D otus
USER otus

COPY --from=builder /app/bin/otusdemo /app/bin/

WORKDIR /app/

EXPOSE 8080
CMD ["./bin/otusdemo"]