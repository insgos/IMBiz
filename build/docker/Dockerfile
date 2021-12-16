FROM debian:bullseye-slim

WORKDIR /app

COPY ./bin/imbizapi .

EXPOSE 18181

ENV GO_ENV=stage

CMD ["/app/imbizapi"]