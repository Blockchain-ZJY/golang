protoc -I . \
	  --go_out=./ \
	    --go_opt=paths=source_relative \
	      ./grpc_proto/hello.proto

protoc -I . \
	  --go-grpc_out=. \
	    --go-grpc_opt=paths=source_relative \
	      ./grpc_proto/hello.proto
