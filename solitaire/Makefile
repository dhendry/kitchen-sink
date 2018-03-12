
build: generate
	go build ./...

#run-webui: build


test: build
	go test ./...

generate: clean-generated
	protoc --go_out=./ model/*.proto

clean-generated:
	find ./ -iname "*.pb.go" -type file -delete
