easyjson:
	easyjson -pkg ./usecase/entity ./adapter/manticore_http

protos:
	#protoc -I./proto  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate --go_out=pb --go_opt=paths=source_relative --validate_out="lang=go:." --go-grpc_out=pb --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative proto/*.proto
	protoc -I./proto  --go_out=pb --go_opt=paths=source_relative  --go-grpc_out=pb --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative proto/*.proto
	protoc-go-inject-tag -input=pb/*.pb.go