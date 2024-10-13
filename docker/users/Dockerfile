FROM golang:1.23

WORKDIR /go/src/app
COPY . .

# install module
RUN go mod download \
    && go install github.com/air-verse/air@latest \

CMD ["air", "-c", ".air.toml"]
