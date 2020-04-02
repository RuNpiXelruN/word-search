.PHONY: proto

proto:
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH} \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc,paths=source_relative:./proto proto/newworsearch.proto  \
		--grpc-gateway_out=./proto \
		--swagger_out=logtostderr=true:./proto

		mv ./proto/proto/* ./proto/
		mv ./proto/github.com/RuNpiXelruN/word-search/* ./proto/
		rm -r ./proto/proto
		rm -r ./proto/github.com

	
	# protoc -I/usr/local/include -I. \
	# 	-I${GOPATH}/src \
	# 	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	# 	--grpc-gateway_out=logtostderr=true:. protowordsearch.proto

	# protoc -I/usr/local/include -I. \
	# 	-I${GOPATH}/src \
	# 	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	# 	--swagger_out=logtostderr=true:. protowordsearch.proto