fmt:
	go fmt ./...

gen:
	./protogen.sh
	./build_info.sh
	go generate ./...

test : gen
	go test ./...

cover : gen
	go test ./... -coverprofile=/tmp/coverage.out && go tool cover -html=/tmp/coverage.out


tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

compile:
	rm build/*
	cd cmd && go build -o ../build -v
