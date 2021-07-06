MODULE := $(shell awk '/^module / {print $$2}' go.mod)

.PHONY: protoc
protoc:
	protoc \
		-I.:proto/lib \
		--go_out=module=$(MODULE):. \
		--go-grpc_out=module=$(MODULE):. \
		--include_imports \
		--include_source_info \
		--descriptor_set_out=api_descriptor.pb \
		proto/api.proto
