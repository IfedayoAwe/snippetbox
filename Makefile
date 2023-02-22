SOURCES := $(wildcard *.go cmd/*/*.go pkg/*/*/*.go)

VERSION=$(shell git describe --tags --long --dirty --always 2>/dev/null)

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

.PHONY: committed
committed:
	@git diff --exit-code > /dev/null || (echo "** COMMIT YOUR CHANGES FIRST **"; exit 1)

docker: $(SOURCES) Dockerfile
	docker build -t snippetbox:latest . -f Dockerfile --build-arg VERSION=$(VERSION)

.PHONY: publish
publish: committed
	make docker
	docker tag snippetbox:latest ifedayoawe/snippetbox:$(VERSION)
	docker push ifedayoawe/snippetbox:$(VERSION)