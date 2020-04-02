.PHONY: proto build test binaries runbinary run

cleanproto: ## Removes all generated proto related files (excluding original .proto file) from ./proto folder
	rm -r ./proto/{*.go,*.json}

cleanbinaries: ## Removes all binaries from ./bin folder
	rm -r ./bin/*

clean: ## Removes all generated proto related files (excluding original .proto file) from ./proto folder, and all binaries from ./bin folder
	rm -r ./proto/{*.go,*.json} && \
	rm -r ./bin/*

proto: ## Generates proto api, grpc-gateway definition, and swagger definition 
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH} \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc,paths=source_relative:./proto proto/wordsearch.proto  \
		--grpc-gateway_out=./proto \
		--swagger_out=logtostderr=true:./proto

		mv ./proto/proto/* ./proto/
		mv ./proto/github.com/RuNpiXelruN/word-search/* ./proto/
		rm -r ./proto/{proto,github.com}

test: ## Runs test suite
	go test ./...

binaries: ## Builds binaries for osx, linux, and windows, and places them in the ./bin folder.
	cd handler && \
	env GOOS=darwin GOARCH=amd64 go build -v -o grpcWordSearch_osx && \
	env GOOS=linux GOARCH=amd64 go build -v -o grpcWordSearch_linux && \
	env GOOS=windows GOARCH=amd64 go build -v -o grpcWordSearch_windows && \
	mv ./*grpcWordSearch* ../bin/

build: test binaries ## Runs test suite, and builds binaries for osx, linux, and windows, and places them in the ./bin folder.
	
run: ## Runs osx version of binary
	cd bin && \
	./grpcWordSearch_osx

buildrun: build run ## Runs tests, builds binaries and runs osx version of binary

generate: proto build run ## Generates proto output, runs tests, builds binaries, and runs osx version of binary
  

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

