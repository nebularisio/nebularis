
REPO_ROOT = $(shell git rev-parse --show-toplevel)

ANTLR_JAR=antlr-4.7.1-complete.jar

COVER_PACKAGES= \
	"nebularis.io/nebularis/pkg/..."

ifeq ($(DIFF),)
	DIFF := opendiff
endif

.PHONY: default
default: check-license build lint test # TODO: Re-add gen target when it can work in CircleCI

.PHONY: gen
gen: lib/$(ANTLR_JAR)
	$(REPO_ROOT)/bin/genparser.sh

lib/$(ANTLR_JAR):
	mkdir -p $(REPO_ROOT)/lib
	curl https://www.antlr.org/download/$(ANTLR_JAR) -o $(REPO_ROOT)/lib/$(ANTLR_JAR)

.PHONY: build
build:
	go build $(REPO_ROOT)/...

test:
	go test -v -race -cover -coverprofile $(REPO_ROOT)/cover.out -coverpkg $(COVER_PACKAGES) $(REPO_ROOT)/tests/...

test-diff:
	rm -rf $(REPO_ROOT)/out
	-go test $(REPO_ROOT)/tests/... -diffout $(REPO_ROOT)/out
	${DIFF} $(REPO_ROOT)/tests/testdata/compiler $(REPO_ROOT)/out -merge $(REPO_ROOT)/tests/testdata/compiler

lint:
	golangci-lint run $(REPO_ROOT)/...

format:
	goimports -local nebularis.io -w .

.PHONY: check-license
check-license:
	$(REPO_ROOT)/bin/checklicense.sh

update-baselines:
	go test -v $(REPO_ROOT)/tests/... -ubl

clean:
	rm -f $(REPO_ROOT)/lib/$(ANTLR_JAR)
	rm -f $(REPO_ROOT)/cover.out
	rm -rf $(REPO_ROOT)/out
