.PHONY: report
report:
	@@go test -v -coverprofile=cover.out $(shell go list ./... | grep -v debug)
	go tool cover -html=cover.out -o cover.html
	open cover.html