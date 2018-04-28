DEP=$(GOPATH)/bin/dep


install: $(DEP)
	$(DEP) ensure

$(DEP):
	go get -u github.com/golang/dep/...

build:
	go build -o trellocli

clean:
	rm -f trellocli

