# TODO:两阶段构建
FROM alpine:latest

WORKDIR /wweb

ADD wweb .

EXPOSE 8080 

ENTRYPOINT ["./wweb", "run"]