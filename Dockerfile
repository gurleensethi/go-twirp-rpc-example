FROM golang:latest

RUN go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

CMD ["protoc --twirp_out=. --go_out=. rpc/notes/service.proto"]