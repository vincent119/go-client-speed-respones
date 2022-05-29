FROM golang:1.18.1-alpine3.15
RUN apk --no-cache add gcc g++ make ca-certificates
RUN apk add git 
RUN apk add tzdata 
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
#RUN useradd -rm -d /home/appUser -s /bin/bash -g root -G sudo -u 1500 appUser
RUN mkdir /go-client-speed-respones
WORKDIR /go-client-speed-respones
RUN addgroup -S appUser
RUN adduser -S -D appUser appUser
RUN chown appUser:appUser /go-client-speed-respones -R
USER appUser
#ADD . /go-client-speed-respones
COPY . .
RUN go env -w GOSUMDB=off
RUN go env -w GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build -o go-client-speed-respones
ENTRYPOINT ["./go-client-speed-respones"]
