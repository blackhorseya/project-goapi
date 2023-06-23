## common
.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## clean artifacts
	@rm -rf ./deployments/charts/*.tgz
	@echo Successfully removed artifacts

.PHONY: version
version: ## show version
	@echo $(VERSION)

.PHONY: gen
gen: gen-pb-go gazelle ## generate all go files
	## Successfully generated all go files

## go
.PHONY: build-go
build-go: ## build go binary
	@bazel build //...

.PHONY: test-go
test-go: ## run go test
	@bazel test //...

## protobuf
.PHONY: gen-pb-go
gen-pb-go: ## generate protobuf messages and services
	@go get -u google.golang.org/protobuf/proto
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=./pb --go_out=paths=source_relative:./entity --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./pb ./pb/domain/*/*/*.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

## bazel
.PHONY: gazelle-repos
gazelle-repos: ## update gazelle repos
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune

.PHONY: gazelle
gazelle: gazelle-repos ## run gazelle with bazel
	@bazel run //:gazelle
