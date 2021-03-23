FROM ubuntu:18.04

COPY ./meta /meta

COPY ./my-app /my-app

RUN chmod +x /my-app

WORKDIR /
ENTRYPOINT my-app