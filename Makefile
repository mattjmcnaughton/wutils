generate_build_files:
	dep ensure && bazel run //:gazelle

build: generate_build_files
	bazel build //pkg/...

clean:
	bazel clean

lint:
	bash -c "! go list ./... | xargs -L1 golint 2>&1 | read"

fmt:
	bash -c "! gofmt -d . 2>&1 | read"

vet:
	go vet ./...

static: lint fmt vet

check: static

correct:
	gofmt -w .
