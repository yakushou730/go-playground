run: generate_go_pb_file

generate_go_pb_file:
	protoc --go_out=. --go_opt=paths=source_relative \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  proto/*.proto
