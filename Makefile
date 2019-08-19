STRINGER_FILES := $(shell find . -name '*.go' -and -not -name '*_string.go' -and -not -iwholename './vendor/**/*' | xargs grep -l '//go:generate stringer')
MOCKGEN_FILES := $(shell find . -name '*.go' -and -not -name '*_mockgen.go' -and -not -iwholename './vendor/**/*' | xargs grep -l -E '//go:generate.*mockgen')
WIRE_FILES := $(shell find . -name '*.go' -and -not -wholename './vendor/**/*' | xargs grep -l '// +build wireinject')
GENERATED_STRINGER_FILES := $(patsubst %.go, %_string.go, $(STRINGER_FILES))
GENERATED_MOCKGEN_FILES := $(patsubst %.go, %_mockgen.go, $(MOCKGEN_FILES))
GENERATED_WIRE_FILES := $(shell find . -name 'wire_gen.go' -and -not -iwholename './vendor/**/*' | xargs grep -l '//go:generate wire')
GENERATED_FILES := $(GENERATED_STRINGER_FILES) $(GENERATED_MOCKGEN_FILES) $(GENERATED_WIRE_FILES)

.PHONY = run build

%.go:
	generate

run:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go

deps:
	# go get -u golang.org/x/tools/cmd/stringer
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	# go get -u github.com/golang/mock/gomock
	# go install github.com/golang/mock/mockgen
	go mod download
	go mod tidy

lint:
	golangci-lint run --fix

generate:
	go generate ./...
	@for f in $(WIRE_FILES); do wire $$f; done

clean:
	@for f in $(GENERATED_FILES); do if [ -f $$f ]; then rm $$f; fi; done
