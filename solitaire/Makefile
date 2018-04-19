.DEFAULT_GOAL := help

build: generate ## Basically just a go build, pluse generate files
	# I THINK this is where dep ensure should be... ???
	dep ensure
	go build ./...

run-webui: generate ## .
	cd webui && go run main.go

test: build ## .
	go test ./...

#generate: clean-generated ## .
generate: ## .
	protoc --go_out=${GOPATH}/src/ model/*.proto
	#protoc -Iwebui/server/ -Imodel/ --go_out=webui/server/ webui/server/*.proto

	#protoc --go_out=${GOPATH}/src/ webui/server/*.proto

	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:${GOPATH}/src/  \
		--grpc-gateway_out=logtostderr=true:${GOPATH}/src/  \
		--swagger_out=logtostderr=true:. \
		webui/server/*.proto
	#protoc --go_out=./ webui/server/*.proto

	#protoc --go_out=./ --js_out=import_style=commonjs,binary:. model/*.proto
	#mv model/*_pb.js ~/code/phaser3-testing/src/

clean-generated: ## .
	find ./ -iname "*.pb.go" -type file -delete
	find ./ -iname "*.pb.gw.go" -type file -delete

	#DONT LIKE THIS:
	find ./ -iname "*.swagger.json" -type file -delete

goget: ## Deprecated in favor of dep
	go get -u github.com/jinzhu/gorm
	#go get -u github.com/jmoiron/sqlx
	go get -u github.com/stretchr/testify
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go


# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
