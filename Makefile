go_out=pb
proto_dir=proto


# require_unimplemented_servers=false 不加这一行GRPC注册服务实例时，提示缺失方法
# --proto_path 指定protobuf的父目录

pb:
	rm -rf $(go_out)/*

	protoc --go_out=$(go_out) \
		   --go-grpc_out=require_unimplemented_servers=false:$(go_out) \
		   --proto_path=$(proto_dir) \
		   $(proto_dir)/*.proto

client:
	cd proto && evans --proto *.proto --port 8080

.PHONY: pb