PROJECT_NAME=imega/teleport-server
GO_PROJECT=github.com/$(PROJECT_NAME)
CWD=/go/src/$(GO_PROJECT)
TAG=latest
IMG=imega/teleport-server

GO_IMG=golang:1.10-alpine

build:
	docker build -t $(IMG):$(TAG) .
	
.PHONY: acceptance
acceptance:
	@touch $(CURDIR)/mysql.log
	@TAG=$(TAG) docker-compose up -d
	@docker run --rm \
		--network teleportserver_default \
		-v $(CURDIR):$(CWD) \
		$(GO_IMG) sh -c "go test -v $(GO_PROJECT)/acceptance"

clean:
	@-rm $(CURDIR)/mysql.log
	@TAG=$(TAG) docker-compose rm -sfv

generate:
	@go generate ./schema

proto:
	@docker run --rm -v $(CURDIR)/api:/data -v $$GOPATH:/go -w /data \
		imega/grpcgen:1.0.0 -I/data -I/go/src/github.com/googleapis/googleapis -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle=github.com/googleapis/googleapis/google,plugins=grpc:. /data/service.proto

test: clean build acceptance
