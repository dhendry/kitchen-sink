
build: generate
	go build ./...

run-webui: build
	cd webui && go run main.go

test: build
	go test ./...

generate: clean-generated
	protoc --go_out=./ model/*.proto
	protoc -Iwebui/server/ -Imodel/ --go_out=webui/server/ webui/server/*.proto
	#protoc --go_out=./ --js_out=import_style=commonjs,binary:. model/*.proto
	#mv model/*_pb.js ~/code/phaser3-testing/src/

clean-generated:
	find ./ -iname "*.pb.go" -type file -delete
