go_out=pb
proto_dir=proto

MODULE=pb/github.com/six1203/order

pb:
	rm -rf $(go_out)/*

	protoc --go_out=$(go_out) \
		   --go-grpc_out=$(go_out) \
		   --proto_path=$(proto_dir) \
		   $(proto_dir)/*.proto

client:
	cd proto && evans --proto *.proto --port 8080

.PHONY: pb