PROTO_PATH := api/proto

# Output directories.
GRPC_OUT := api/generated

protoc:
	# Generate proto stubs.
	protoc \
	-I $(PROTO_PATH) \
	--go_out=plugins=grpc:$(GRPC_OUT) \
	$(PROTO_PATH)/*.proto
