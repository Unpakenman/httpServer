GOLANGCI_BIN=$(LOCAL_BIN)/golangci-lint
$(GOLANGCI_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

MOCKERY_BIN=$(LOCAL_BIN)/mockery
$(MOCKERY_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.20.0
