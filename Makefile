# Setup name variables for the package/tool
NAME := passwordgetter
PKG := github.com/petems/$(NAME)

# Set the build dir, where built cross-compiled binaries will be output
BUILDDIR := ${PREFIX}/builds

.PHONY: all
all: clean build fmt lint test install

.PHONY: build
build: $(NAME) ## Builds a dynamic executable or package

$(NAME):
	go build -o $(NAME) .

.PHONY: fmt
fmt: ## Verifies all files have men `gofmt`ed
	@echo "+ $@"
	@gofmt -s -l . | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint: ## Verifies `golint` passes
	@echo "+ $@"
	@golint ./... | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

.PHONY: cover
cover: ## Runs go test with coverage
	@for d in $(shell go list ./... | grep -v vendor | grep -v mocks); do \
		go test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
	done;

.PHONY: cover_html
cover_html: cover ## Opens html of coverage
	@go tool cover -html=profile.out

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	$(RM) $(NAME)
	$(RM) -r $(BUILDDIR)

.PHONY: test
test: ## Runs the go tests
	@echo "+ $@"
	@go test ./...

.PHONY: generate
generate: ## runs generate for mocks
	@go generate ./...

.PHONY: install
install: ## Installs the executable or package
	@echo "+ $@"
	go install -a .
