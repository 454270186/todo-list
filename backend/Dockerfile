FROM golang:1.20

ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /app

ADD . /app/

EXPOSE 8080

RUN go build -o main ./cmd

CMD [ "./main" ]