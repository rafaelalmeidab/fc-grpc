# gRPC com Go — Full Cycle

Projeto de estudos do curso de gRPC da Full Cycle, usando Go, Protocol Buffers e o exemplo `helloworld`.

## 🐳 Ambiente de desenvolvimento (Docker)

O projeto usa um container de desenvolvimento (Go + `protoc` + plugins já instalados), então **não precisa instalar nada localmente** — só Docker.

### Subindo o container

```bash
docker-compose up -d
```

Isso builda a imagem (a partir do `Dockerfile`) e sobe o container `go-app`, com o código montado como volume (`.:/app`) e a porta `8282` exposta.

### Acessando o container

Como o container fica apenas em segundo plano (`sleep infinity`), os comandos abaixo (proto, build, run) devem ser executados **dentro dele**:

```bash
docker exec -it go-app bash
```

### Parando o container

```bash
docker-compose down
```

## 🔧 Gerando o código a partir do `.proto`

Dentro do container:

```bash
protoc --go_out=. --go_opt=module=github.com/rafaelalmeidab/fc-grpc \
  --go-grpc_out=. --go-grpc_opt=module=github.com/rafaelalmeidab/fc-grpc \
  examples/helloworld/proto/helloworld.proto
```

Esse comando lê o `helloworld.proto` e gera os arquivos Go correspondentes (`*.pb.go` e `*_grpc.pb.go`), usados para serializar mensagens e implementar o servidor/cliente gRPC.

## 📚 Go Modules — para que serve?

O `go.mod` gerencia as dependências do projeto (parecido com `package.json` no Node ou `composer.json` no PHP).

```bash
go mod tidy
```

Adiciona ao `go.mod`/`go.sum` os pacotes que estão sendo importados no código e remove os que não são mais usados.

## 🏗️ Build

```bash
go build -o bin/helloworld/server ./helloworld/server
```

Gera o binário do servidor em `bin/helloworld/server`.

## ▶️ Rodando o servidor

```bash
./bin/helloworld/server
```

O servidor gRPC sobe e fica no ar, escutando na porta configurada (por padrão exposta em `8282` via `docker-compose.yml`).

## 📂 Estrutura do projeto

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
