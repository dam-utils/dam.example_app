FROM ubuntu:18.04

COPY ./meta /meta

COPY ./my-app /opt/my-app

RUN chmod +x /opt/my-app

ENTRYPOINT /opt/my-app