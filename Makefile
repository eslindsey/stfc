.PHONY: all
all: *.pb.go *.proto
	go test -v

%.pb.go: %.proto
	protoc --experimental_allow_proto3_optional -I. --go_out=. --go_opt=paths=source_relative $^

