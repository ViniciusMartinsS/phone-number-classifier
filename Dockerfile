FROM golang:latest

ENV GO111MODULE "on"

RUN mkdir /app

WORKDIR /app
COPY . /app

RUN cd /app && go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags="-s -w" -o /phone-number-classifier

EXPOSE 8080

CMD [ "/phone-number-classifier" ]
