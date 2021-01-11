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

.PHONY: download-mod
download-mod:
	@go mod download