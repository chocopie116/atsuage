setup:
	go get -u github.com/golang/dep/cmd/dep

test:
	go test -v $$(go list ./... | grep -v /vendor/)
