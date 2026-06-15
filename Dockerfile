FROM golang:1.25

RUN apt-get update && \
    apt-get install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="${PATH}:/root/go/bin"

WORKDIR /app

COPY . .

CMD ["sleep", "infinity"]