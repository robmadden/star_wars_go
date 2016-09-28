FROM golang:1.7

ADD . /go/src/github.com/spbk/star_wars_go

RUN go get github.com/graphql-go/graphql
RUN go get github.com/graphql-go/handler
RUN go get github.com/graphql-go/relay
RUN go get github.com/object88/relay/examples/starwars
RUN go install github.com/spbk/star_wars_go

ENTRYPOINT /go/bin/star_wars_go

EXPOSE 3000
