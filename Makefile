DEP=$(GOPATH)/bin/dep


install: $(DEP)
	$(DEP) ensure

$(DEP):
	go get -u github.com/golang/dep/...

setup: 
	cp config.toml.sample config.toml
