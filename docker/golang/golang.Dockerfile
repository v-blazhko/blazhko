FROM golang:alpine
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PORT=3001
WORKDIR /build
COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download
COPY ./backend .
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
EXPOSE 3000
CMD ["/dist/main"]