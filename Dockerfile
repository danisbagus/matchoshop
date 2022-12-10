FROM golang:1.15-alpine  as builder
RUN apk add build-base
RUN mkdir /app 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go mod vendor
RUN go build main.go
RUN cd app/migration && go build main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder app/app/migration/ /root/app/migration/
COPY --from=builder app/main /root/main
CMD ["sh", "-c", "/root/app/migration/main up & /root/main"]