GOCMD=go
GOTEST=$(GOCMD) test
GORUN=${GOCMD} run

start: 
		${GORUN} ./cmd/webserver/main.go