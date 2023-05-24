.PHONY: all
all: sync.pb.go
	go test -v

%.pb.go: %.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative $^

