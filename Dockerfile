FROM alpine:latest

RUN apk add go
RUN apk add rsync

COPY . .

RUN mkdir /from
RUN mkdir /to

RUN go build -o /backup main.go

CMD ["/backup"]
