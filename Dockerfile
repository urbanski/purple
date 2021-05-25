# syntax=docker/dockerfile:1
FROM golang:1.16
WORKDIR /go/src/github.com/urbanski/purple
COPY . /go/src/github.com/urbanski/purple/
RUN go build -o purple .
RUN cp purple /root
RUN chmod +x /root/purple
CMD ["/root/purple","-h"]