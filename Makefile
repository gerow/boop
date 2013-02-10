all:
	@export OLDGOPATH=$(GOPATH) && \
	export GOPATH=`pwd` && \
	go install boop/boop && \
	export GOPATH=$(OLDGOPATH)

test:
	@export OLDGOPATH=$(GOPATH) && \
	export GOPATH=`pwd` && \
	go get -u github.com/gerow/gotest && \
	go test boop && \
	export GOPATH=$(OLDGOPATH)

clean:
	rm -rf pkg bin
