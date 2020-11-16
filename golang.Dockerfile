FROM golang:alpine as golang
ENV GOPATH=/go/src/app
ADD ./backend /go/src/app
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

WORKDIR /go/src/app/
#RUN go get .
#RUN go install
#RUN go build main.go

#RUN oapi-codegen -generate spec -package api openapi.yaml > api/spec.go
#    oapi-codegen -generate types -package api openapi.yaml > api/types.go
#    oapi-codegen -generate chi-server -package api openapi.yaml > api/server.go

ENV PORT=3001
CMD ["go", "run", "main.go"]