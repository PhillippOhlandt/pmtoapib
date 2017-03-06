FROM golang:1.7.5

MAINTAINER 	Phillipp Ohlandt <phillipp.ohlandt@googlemail.com>

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go build -o pmtoapib . && \
    mv /app/pmtoapib /usr/local/bin/pmtoapib && \
    chmod +x /usr/local/bin/pmtoapib

WORKDIR /opt

ENTRYPOINT ["pmtoapib"]
