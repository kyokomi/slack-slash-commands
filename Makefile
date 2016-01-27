ROOT_GOPATH=$(shell pwd)
VENDOR_GOPATH=$(shell pwd)/vendor
GOPATH=$(VENDOR_GOPATH)

install:
	mkdir -p $(VENDOR_GOPATH)
	sh ./goapp.sh

clean:
	rm -rf $(VENDOR_GOPATH)

run:
	goapp serve ./src

clean-run:
	goapp serve --clear_datastore ./src

test:
	GOPATH=$(VENDOR_GOPATH):$(ROOT_GOPATH) goapp test -v ./src/...

deploy:
	goapp deploy ./src


