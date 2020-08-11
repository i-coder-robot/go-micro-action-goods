.PHONY: proto
proto:
	protoc -I . --micro_out=. --gogofaster_out=. proto/category/category.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/product/product.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/carousel/carousel.proto
