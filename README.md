https://grpc.io/docs/languages/go/basics/
https://protobuf.dev/reference/go/go-generated/
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ./proto/articles_service.proto


https://earthly.dev/blog/golang-sqlite/
sqlite
brew install sqlite3
brew install sqlite-utils