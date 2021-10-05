FROM golang:1.17.1-buster

WORKDIR /release
COPY . .
RUN go get -v -d ./...
RUN CGO_ENABLED=0 go build -a -o eCommerce ./main.go

FROM alpine:3.14
WORKDIR /root/

COPY --from=0 /release/products.json .
COPY --from=0 /release/eCommerce .
COPY --from=0 /release/.env .
CMD ["./eCommerce"]