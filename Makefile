# list everything except for vendored pkgs, then lint, test, and install
CMD_DIR := cmd/goprimer.go
PKGS := $(shell go list github.com/DanielSchuette/cloningprimer/ | grep -v /vendor)
EXEC := $(BIN_DIR)/goprimer.go
BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(EXEC): test windows
	go install $(CMD_DIR)

.PHONY: test
test: lint 
	go test -v -cover $(PKGS)

# go version 1.11 broke gometalinter's linter
# super annoying but I don't have a fix at this point
.PHONY: lint
lint: $(GOMETALINTER)
	# gometalinter . --enable=gofmt --enable=gosimple --enable=staticcheck --disable=gocyclo --vendor
	echo "no linting at this point"

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

# create binaries for windows, linux, and darwin and put them inside a newly created bin/
LOCAL_BIN_DIR := bin/goprimer
VERSION = 0.0.3

.PHONY: windows
windows: linux darwin 
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o $(LOCAL_BIN_DIR)-v$(VERSION)-windows-amd64 $(CMD_DIR)

.PHONY: linux
linux: darwin 
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o $(LOCAL_BIN_DIR)-v$(VERSION)-linux-amd64 $(CMD_DIR)

.PHONY: darwin
darwin: 
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -o $(LOCAL_BIN_DIR)-v$(VERSION)-darwin-amd64 $(CMD_DIR)
