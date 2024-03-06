.PHONY: all

proto_gen1:
	@echo "Hello proto build........."
#	原地生产代码
	protoc --proto_path=. \
		   --go_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
		   pg/*.proto
proto_gen2:
	@echo "Hello proto build........."
	protoc -I ./pb \
			--go_out=hello pb/*.proto\
			--go-grpc_out=./hello pb/*.proto
