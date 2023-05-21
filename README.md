# Go GRPC server

Sample of GRPC server written on go

## Setup protobuf

You can skip this section if you do not want to
change [existing protobuf](https://github.com/Nikita-Filonov/sample_go_grpc_server/blob/main/proto/articles_service.proto)

1. Install protobuf on your system
    - [MacOS/Linux](https://grpc.io/docs/protoc-installation/)
    - [Windows](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/)
2. Install `protoc-gen-go-grpc` for go `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`. This is used
   to generate client/server code for golang

Generate client, server golang code with command:

```shell
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ./proto/my_service.proto
```

## Setup project

You have to install [go](https://go.dev/doc/install)
and [sqlite3](https://www.tutorialspoint.com/sqlite/sqlite_installation.htm)

Clone the project and install go packages

```shell
git clone https://github.com/Nikita-Filonov/sample_go_grpc_server`
cd ./sample_go_grpc_server
go mod download
```

Apply migrations

```shell
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up -config=./infrastructure/config.yml
```

Finally run the server

```shell
go mod download
```