run: clean generate_go_pb_file

generate_go_pb_file:
	protoc -I ./proto \
	  --go_out=./proto \
	  --go_opt paths=source_relative \
	  --go-grpc_out=./proto \
	  --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out=logtostderr=true:./proto \
	  --grpc-gateway_opt paths=source_relative \
	  ./proto/*.proto

clean:
	rm -f ./proto/*.pb.go ./proto/*.pb.gw.go
