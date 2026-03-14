.DEFAULT_GOAL := help
LOCAL_BIN=$(CURDIR)/bin

include bin-deps.mk

.PHONY: getnewprotovervion
run: ## generate proto code project
	$ go get github.com/Unpakenman/protos/gen/go/sso@v0.1.2





