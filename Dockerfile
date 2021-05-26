# syntax=docker/dockerfile:1
FROM golang:1.16
WORKDIR /go/src/github.com/urbanski/purple
COPY . /go/src/github.com/urbanski/purple/
RUN go build -o purplecli .
RUN cp purplecli /root
RUN chmod +x /root/purplecli
CMD ["/root/purplecli", "-h"]