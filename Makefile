.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: install
# install protoc-gen-go-kit
install:
	go install

.PHONY: prototest
# generate proto test code
prototest:
	cd ./example && protoc --proto_path=. \
		--proto_path=../third_party \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--go-kit_out=paths=source_relative:. \
		echo.proto
