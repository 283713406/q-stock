FROM alpine

WORKDIR /srv/q-stock

ADD ./dist/q-stock.tar.gz /srv/

EXPOSE 4869 4870
ENTRYPOINT ["./q-stock", "-c", "./config.toml"]
