slack-slash-commands
=================

## Setup

Edit the "install:" target of the Makefile.

```
install:
	mkdir -p $(VENDOR_GOPATH)
	gae get <dependencies package>
```

```
$ make install
```

## Run

```
$ make run
```

## Test

```
$ make test
```

## Deploy

```
$ make deploy
```
