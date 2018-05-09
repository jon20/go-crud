FROM golang

WORKDIR /go/src/app

EXPOSE 9000

ADD  .  /go/src/

CMD ["/usr/local/go/bin/go", "run", "/go/src/main.go"]
RUN go get -u  gopkg.in/mgo.v2
RUN go get -u  gopkg.in/mgo.v2/bson
RUN go get -u  github.com/gorilla/mux
RUN go get -u   github.com/BurntSushi/toml
