FROM golang:1.18.1-alpine3.15
RUN apk --no-cache add gcc g++ make ca-certificates
RUN apk add git 
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
RUN mkdir /go-client-speed-respones
WORKDIR /go-client-speed-respones
#ADD . /go-client-speed-respones
COPY . .
RUN go env -w GOSUMDB=off
RUN go env -w GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build -o go-client-speed-respones
ENTRYPOINT ["./go-client-speed-respones"]
