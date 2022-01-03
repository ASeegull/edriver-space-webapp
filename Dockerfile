# Read about multistage images
FROM golang:1.17-alpine as builder

WORKDIR /src

COPY . .

# build stage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o webapp cmd/app/main.go

# actual image
FROM alpine:3.14

LABEL GROUP Lv-644.Golang

RUN apk add --update --no-cache ca-certificates
WORKDIR /usr/lib/edriverspace
COPY --from=builder /src/webapp /usr/lib/edriver-space/webapp
RUN chmod +x /usr/lib/edriver-space/webapp

ENTRYPOINT [ "/usr/lib/edriver-space/webapp" ]

EXPOSE 5050
