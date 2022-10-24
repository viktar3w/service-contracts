.PHONY: proto
proto: clean format gen lint

# =========== BUF ===========
BUF_VERSION=v1.6.0

.PHONY: buf-install
buf-install:
	echo $(GOPATH)

gen: buf-install
	@$(GOPATH)/bin/buf generate
	@for dir in $(CURDIR)/gen/go/*/; do \
  	  cd $$dir && \
  	  go mod init && go mod tidy; \
  	done

lint: buf-install
	@$(GOPATH)/bin/buf lint

format: buf-install
	@$(GOPATH)/bin/buf format

clean:
	@rm -rf ./gen || true