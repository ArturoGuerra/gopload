.PHONY: all build clean docker-build docker-push docker docker-check

APPNAME = goimgupload
GOBUILD = go build
DOCKER  = docker

all: clean build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/*.go

docker-check:
	test $(DOCKERREPO)

docker-build: docker-check
	$(DOCKERREPO) build . -t $(DOCKERREPO)

docker-push: docker-check
	$(DOCKERREPO) push $(DOCKERREPO)

docker: docker-build docker-push
