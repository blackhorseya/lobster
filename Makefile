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

.PHONY: build-image
build-image:
	@docker build -t $(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	-f Dockerfile .

.PHONY: gen
gen: gen-wire

.PHONY: gen-wire
gen-wire:
	@wire gen ./...
