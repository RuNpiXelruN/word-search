.PHONY: proto

clean:
	rm -r ./proto/{*.go,*.json}

proto:
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