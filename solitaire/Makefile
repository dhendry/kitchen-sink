
build: generate
	go build ./...

run-webui: build
	cd webui && go run main.go

test: build
	go test ./...

generate:
	protoc --go_out=./ model/*.proto
	#protoc --go_out=./ --js_out=import_style=commonjs,binary:. model/*.proto
	#mv model/*_pb.js ~/code/phaser3-testing/src/

clean-generated:
	find ./ -iname "*.pb.go" -type file -delete
