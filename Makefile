proto:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		./selector/selector.proto

test:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-kratos-selector_out=paths=source_relative:. \
		./tests/protos/certs.proto