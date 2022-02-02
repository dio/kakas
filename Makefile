# Copyright 2022 Dhi Aurrahman
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: test

# Include versions of tools we build or fetch on-demand.
include Tools.mk

root_dir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Local cache directory.
CACHE_DIR ?= $(root_dir).cache

go_tools_dir := $(CACHE_DIR)/tools/go

# Go-based tools.
addlicense    := $(go_tools_dir)/addlicense
buf           := $(go_tools_dir)/buf
golangci-lint := $(go_tools_dir)/golangci-lint
protoc-gen-go := $(go_tools_dir)/protoc-gen-go

proto_tools := \
	$(buf) \
	$(protoc-gen-go)

export PATH := $(root_dir)test:$(PATH)
test: $(proto_tools)
	@$(call go-build,protoc-gen-deepcopy)
	@$(call go-build,protoc-gen-jsonshim)
	cd $(root_dir)test && buf generate && go test ./...

define go-build
	go build -o $(root_dir)test/$1 ./$1
endef

license_ignore :=
license_files  := test Makefile Tools.mk protoc-gen-deepcopy protoc-gen-jsonshim
license: $(addlicense) ## To add license
	@$(addlicense) $(license_ignore) -c "Dhi Aurrahman"  $(license_files) 1>/dev/null 2>&1

# Override lint cache directory. https://golangci-lint.run/usage/configuration/#cache.
export GOLANGCI_LINT_CACHE=$(CACHE_DIR)/golangci-lint
lint: .golangci.yml $(golangci-lint) ## Lint all Go sources
	@$(golangci-lint) run --timeout 5m --config $< ./...

check: lint license

$(go_tools_dir)/%:
	@GOBIN=$(go_tools_dir) go install $($(notdir $@)@v)
