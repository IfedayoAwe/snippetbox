SOURCES := $(wildcard *.go cmd/*/*.go pkg/*/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

docker: $(SOURCES) build/Dockerfile
	docker build -t snippetbox:latest . -f build/Dockerfile --build-arg VERSION=$(VERSION)

.PHONY: publish
publish:
	make docker
	docker tag  sort-anim:latest matthol2/sort-anim:$(VERSION)
	docker push matthol2/sort-anim:$(VERSION)