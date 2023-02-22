SOURCES := $(wildcard *.go cmd/*/*.go pkg/*/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

.PHONY: committed
committed:
	@git diff --exit-code > /dev/null || (echo "** COMMIT YOUR CHANGES FIRST **"; exit 1)

docker: $(SOURCES) build/Dockerfile
	docker build -t snippetbox:latest . -f build/Dockerfile --build-arg VERSION=$(VERSION)

.PHONY: publish
publish: committed
	make docker
	docker tag  sort-anim:latest matthol2/sort-anim:$(VERSION)
	docker push matthol2/sort-anim:$(VERSION)