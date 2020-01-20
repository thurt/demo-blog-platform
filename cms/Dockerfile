FROM golang:1.9 as builder
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download
RUN CGO_ENABLED=0 go-wrapper install

FROM alpine:latest
WORKDIR /root/
RUN apk -v --no-cache add bash ca-certificates coreutils
COPY --from=builder /go/src/app/cms.swagger.json /go/src/app/wait-for-it.sh /go/bin/cms ./
RUN mkdir migrations
COPY --from=builder /go/src/app/mysqlprovider/db/migrations/* ./migrations/
CMD ["/root/cms"]
EXPOSE 10000/TCP 8080/TCP 8181/TCP
