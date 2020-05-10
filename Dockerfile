FROM alpine:3.11.5
RUN adduser -D otus
USER otus

WORKDIR /app/
COPY ./bin/otusdemo /app/bin/

EXPOSE 8080
CMD ["./bin/otusdemo"]