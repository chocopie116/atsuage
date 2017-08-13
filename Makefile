setup:
	go get -u github.com/golang/dep/cmd/dep
	ls /go/bin

install: setup
	dep ensure

test:
	go test -v $$(go list ./... | grep -v /vendor/)
