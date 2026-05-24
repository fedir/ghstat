.PHONY: build test vet lint clean cache-clear clone-clear rate-limit run-go run-go-microservices run-all help

BINARY := ghstat
STATS_DIR := stats
CACHE_DIR := tmp

## build: compile the binary
build:
	go build -o $(BINARY) .

## test: run tests with race detector and coverage
test:
	@echo "" > coverage.txt
	@for d in $$(go list ./...); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		if [ -f profile.out ]; then cat profile.out >> coverage.txt && rm profile.out; fi \
	done

## vet: run go vet
vet:
	go vet ./...

## clean: remove binary and API cache (preserves local clones in tmp/projects/)
clean:
	rm -f $(BINARY)
	rm -rf $(CACHE_DIR)

## cache-clear: clear HTTP response cache (preserves local clones)
cache-clear: build
	./$(BINARY) -cc -t $(CACHE_DIR)

## clone-clear: remove locally cloned repositories (tmp/projects/ can be several GB)
clone-clear:
	rm -rf $(CACHE_DIR)/projects

## rate-limit: show current GitHub API rate limit status
rate-limit: build
	./$(BINARY) -l

## run-go: fetch and rank Go frameworks
run-go: build
	./$(BINARY) -f $(STATS_DIR)/go_frameworks.csv -t $(CACHE_DIR)

## run-go-microservices: fetch and rank Go microservice toolkits
run-go-microservices: build
	./$(BINARY) -r koding/kite,nytimes/gizmo,micro/go-micro,rsms/gotalk,gocircuit/circuit,go-kit/kit \
		-f $(STATS_DIR)/go_microservice_toolkits.csv -t $(CACHE_DIR)

## run-all: run all framework/CMS comparisons
run-all: build
	bash bin/build_all.sh

## help: show this help
help:
	@grep -E '^## ' Makefile | sed 's/^## //' | column -t -s ':'
