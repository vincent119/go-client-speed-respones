FROM golang:1.18.1-alpine3.15 
RUN apk --no-cache add gcc g++ make ca-certificates
RUN apk add git
WORKDIR /go-client-speed-respones
#ADD . /go-client-speed-respones
COPY . .
RUN go mod download
RUN go build -o go-client-speed-respones
ENTRYPOINT ["./go-client-speed-respones"]
