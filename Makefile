ROOT_GOPATH=$(shell pwd)
VENDOR_GOPATH=$(shell pwd)/vendor
GOPATH=$(VENDOR_GOPATH)

install:
	mkdir -p $(VENDOR_GOPATH)
	-goapp get -u google.golang.org/appengine
	goapp get -u golang.org/x/net/context
	goapp get -u github.com/unrolled/render

clean:
	rm -rf $(VENDOR_GOPATH)

run:
	goapp serve ./src

test:
	GOPATH=$(VENDOR_GOPATH):$(ROOT_GOPATH) goapp test -v ./src/...

deploy: test
	goapp deploy ./src

