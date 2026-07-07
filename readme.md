# gRPC with Go — Full Cycle

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=flat&logo=googlecloud&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
![Full Cycle](https://img.shields.io/badge/Full%20Cycle-Course-orange?style=flat)

> gRPC study project in Go — [Full Cycle](https://fullcycle.com.br/) course, using Protocol Buffers and the classic `helloworld` example.

## 🐳 Development environment (Docker)

The project uses a development container (Go + `protoc` + plugins already installed), so **there's no need to install anything locally** — just Docker.

### Starting the container

```bash
docker-compose up -d
```

This builds the image (from the `Dockerfile`) and starts the `go-app` container, with the code mounted as a volume (`.:/app`) and port `8282` exposed.

### Accessing the container

Since the container just runs in the background (`sleep infinity`), the commands below (proto, build, run) must be executed **inside it**:

```bash
docker exec -it go-app bash
```

### Stopping the container

```bash
docker-compose down
```

## 🔧 Generating code from the `.proto` file

Inside the container:

```bash
protoc --go_out=. --go_opt=module=github.com/rafaelalmeidab/fc-grpc \
  --go-grpc_out=. --go-grpc_opt=module=github.com/rafaelalmeidab/fc-grpc \
  examples/helloworld/proto/helloworld.proto
```

This command reads `helloworld.proto` and generates the corresponding Go files (`*.pb.go` and `*_grpc.pb.go`), used to serialize messages and implement the gRPC server/client.

## 📚 Go Modules — what's it for?

`go.mod` manages the project's dependencies (similar to `package.json` in Node or `composer.json` in PHP).

```bash
go mod tidy
```

Adds the packages being imported in the code to `go.mod`/`go.sum` and removes the ones that are no longer used.

## 🏗️ Build

```bash
go build -o bin/helloworld/server ./helloworld/server
```

Generates the server binary at `bin/helloworld/server`.

## ▶️ Running the server

```bash
./bin/helloworld/server
```

The gRPC server starts and stays up, listening on the configured port (exposed by default on `8282` via `docker-compose.yml`).

## 📂 Project structure

```
.
├── bin/helloworld/
├── examples/helloworld/
│   └── proto/
│       └── helloworld.proto
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── readme.md
```
