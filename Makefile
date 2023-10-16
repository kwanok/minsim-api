.PHONY: generate-post-proto
generate-post-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		server/post/post.proto


.PHONY: generate-minsim-proto
generate-minsim-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		server/minsim/minsim.proto

.PHONY: generate-comment-proto
generate-comment-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		server/comment/comment.proto