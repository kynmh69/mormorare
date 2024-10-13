FROM golang:1.23

WORKDIR /go/src/app
COPY ./go/src .

# install module
RUN go mod download \
    && go install github.com/air-verse/air@latest; \
    go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /go/src/app/cmd/users

CMD ["air", "-c", ".air.toml"]
