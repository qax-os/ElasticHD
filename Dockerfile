FROM alpine:latest

RUN apk add --no-cache curl && \
    cd /tmp && \
    curl -L https://github.com/farmerx/ElasticHD/releases/download/1.2/elasticHD_linux_amd64.zip > elasticHD_linux_amd64.zip && \
    unzip elasticHD_linux_amd64.zip -d /usr/local/bin && \
    rm -f elasticHD_linux_amd64.zip

RUN adduser -D elastic

USER elastic

EXPOSE 9800

ENTRYPOINT ["ElasticHD"]
