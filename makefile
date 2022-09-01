
SHA1				:= $(shell git rev-parse --verify HEAD)
SHA1_SHORT			:= $(shell git rev-parse --verify --short HEAD)
VERSION				:= $(shell cat VERSION.txt)
INTERNAL_BUILD_ID	:= $(shell [ -z "${GITHUB_RUN_NUMBER}" ] && echo "0" || echo ${GITHUB_RUN_NUMBER})
PWD					:= $(shell pwd)
VERSION_HASH		:= ${VERSION}.${INTERNAL_BUILD_ID}

BUILD_IMAGE			:= golang:1.19
LINT_IMAGE			:= golangci/golangci-lint:v1.36.0
SHELL_LINT_IMAGE	:= koalaman/shellcheck:v0.7.1
MARKDOWN_LINT_IMAGE := tmknom/markdownlint

ENVIRONMENT 		?= local

.DEFAULT_GOAL 		:= test

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: version ## Generates the BUILD_VERSION.txt
version:
	@echo "Setting build to Version: v$(VERSION)"
	$(shell echo v$(VERSION_HASH) > BUILD_VERSION.txt)

.PHONY: fmt
fmt: version ## Runs `go fmt` within a docker container
	@echo "Running fmt"

	docker run --rm \
	-v $(PWD):/usr/src/app \
	-w /usr/src/app $(BUILD_IMAGE) \
	go fmt ./...

	@echo "Completed fmt"

.PHONY: lint
lint: version ## Runs more than 60 different linters using golangci-lint, shellcheck and markdownlint to ensure consistency in code.
	@echo "Running Lint"

	docker run --rm \
	-e GOPACKAGESPRINTGOLISTERRORS=1 \
	-e GO111MODULE=on \
	-v $(PWD):/usr/src/app \
	-w /usr/src/app \
	$(LINT_IMAGE) \
	golangci-lint run --timeout=2m

	find . -name "*.sh" | xargs docker run --rm \
	-v $(PWD):/usr/src/app \
	-w /usr/src/app \
	$(SHELL_LINT_IMAGE)

	docker run -i --rm \
	-v $(PWD):/work \
	$(MARKDOWN_LINT_IMAGE)

	@echo "Completed Lint"

.PHONY: test
test: version ## Runs `go test` within a docker container
	@echo "Running Tests"

	docker run --rm \
	-v $(PWD):/usr/src/app \
	-w /usr/src/app $(BUILD_IMAGE) \
	go test -cover -race -coverprofile=coverage.txt -v -p 8 -count=1 ./...

	@echo "Completed tests"
