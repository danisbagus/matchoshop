FROM golang:latest as builder
RUN mkdir /app 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/pressly/goose/cmd/goose
COPY . .
RUN go mod vendor
RUN rm .env
RUN mv .env.production .env
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o bin/matchoshop  .
# ENTRYPOINT ["./bin/matchoshop"]