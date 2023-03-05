go_out=pb
proto_dir=proto

pb:
	rm -rf $(go_out)/*

	protoc --go_out=$(go_out) --go-grpc_out=$(go_out) --proto_path=$(proto_dir) $(proto_dir)/*.proto

client:
	evans --proto $(proto_dir)/*.proto --port 8080

.PHONY: pb