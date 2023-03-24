FROM golang:1.15-alpine  as builder
RUN apk add build-base
RUN mkdir /app 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/steinbacher/goose/cmd/goose
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

FROM scratch
COPY --from=builder /main ./
ENTRYPOINT ["./main"]