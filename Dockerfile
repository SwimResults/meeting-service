# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

WORKDIR /app

COPY meeting-service /app/service
RUN chmod +x /app/service

ENV SR_MEETING_PORT=8080

RUN apk add tzdata
RUN ln -s /usr/share/zoneinfo/Europe/Berlin /etc/localtime

EXPOSE 8080

ENTRYPOINT [ "./service" ]
