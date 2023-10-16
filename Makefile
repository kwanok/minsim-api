.PHONY: generate-reddit-submission-proto
generate-reddit-submission-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		service/reddit/submission/submission.proto


.PHONY: generate-minsim-proto
generate-minsim-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		service/minsim/minsim.proto

.PHONY: generate-reddit-comment-proto
generate-reddit-comment-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		service/reddit/comment/comment.proto