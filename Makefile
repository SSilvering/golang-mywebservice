GOCMD=go
GOTEST=$(GOCMD) test
GORUN=${GOCMD} run

.PHONY: start
start: 
		${GORUN} ./cmd/webserver/main.go