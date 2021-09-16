FROM ubuntu:18.04

COPY ./meta /meta

COPY ./my-app /opt/my-app

RUN chmod +x /opt/my-app

ENTRYPOINT /opt/my-app
ENV DAM_APP_FAMILY=my-app-develop
ENV DAM_APP_SERVERS=localhost:5000,