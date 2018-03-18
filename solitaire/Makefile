
build: generate
	go build ./...

run-webui:
	cd webui && go run main.go

test: build
	go test ./...

generate: clean-generated
	protoc --go_out=./ model/*.proto

clean-generated:
	find ./ -iname "*.pb.go" -type file -delete
