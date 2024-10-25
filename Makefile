include .env
LOCAL_BIN:=$(CURDIR)/bin # немедленное присваивание, а не при исполнении

install-deps: # установка бинарных файлов (плагинов)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-topic-api:
	mkdir -p pkg/topic_v1
	protoc --proto_path api/topic_v1 \
	--go_out=pkg/topic_v1 \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go_opt=paths=source_relative \
	--go-grpc_out=pkg/topic_v1 \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--go-grpc_opt=paths=source_relative  \
	api/topic_v1/topic.proto

run:
	go run cmd/app/main.go

