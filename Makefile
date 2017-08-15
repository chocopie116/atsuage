REFLEX=$(shell go env GOPATH)/bin/reflex
RICHGO=$(shell go env GOPATH)/bin/richgo

setup:
	go get -u github.com/golang/dep/cmd/dep

install: setup
	dep ensure

test: $(RICHGO)
	$(RICHGO) test -v $$(go list ./... | grep -v /vendor/)


test/watch: $(REFLEX)
	$(REFLEX) --regex='\.go$$' --inverse-regex='vendor/*' -- $(MAKE) test

$(REFLEX):
	go get -u github.com/cespare/reflex

$(RICHGO):
	go get -u github.com/kyoh86/richgo

