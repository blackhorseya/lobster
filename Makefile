APP_NAME = lobster
VERSION = latest
PROJECT_ID = sean-side
NS = side
DEPLOY_TO = uat

.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: lint
lint:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/lobster'

.PHONY: test-with-coverage
test-with-coverage:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: download-mod
download-mod:
	@go mod download