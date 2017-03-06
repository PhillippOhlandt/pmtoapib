FROM alpine:3.5

MAINTAINER 	Phillipp Ohlandt <phillipp.ohlandt@googlemail.com>

COPY pmtoapib-linux /usr/local/bin/pmtoapib
RUN chmod +x /usr/local/bin/pmtoapib

WORKDIR /opt

ENTRYPOINT ["pmtoapib"]
