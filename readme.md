## Comando proto

protoc \
  --go_out=. \
  --go_opt=module=github.com/rafaelalmeidab/fc-grpc \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/rafaelalmeidab/fc-grpc \
  examples/helloworld/proto/helloworld.proto

## Go Mod - Para que serve?

go mod tidy

## Build

go build -o bin/helloworld/server ./helloworld/server

## Server no AR

 ./bin/helloworld/server