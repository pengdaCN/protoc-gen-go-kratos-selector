proto:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		./selector/selector.proto

test:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-kratos-selectorV1_out=paths=source_relative:. \
		./tests/protos/certs.proto ./tests/protos/middleware.proto
