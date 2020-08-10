.PHONY: proto
proto:
	protoc -I . --micro_out=. --gogofaster_out=. proto/category/category.proto
