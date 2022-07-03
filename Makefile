TINYGO_DOCKER ?= docker run --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) tinygo/tinygo:0.23.0 tinygo
TINYGO ?= tinygo
UPX ?= upx --brute
STRIP ?= strip

OUTDIR ?= $(CURDIR)/out
APPNAME ?= $(shell basename $(CURDIR))
SOURCES ?= 	$(wildcard *.go) \
			$(wildcard internal/nostd/*.go) \
			$(wildcard internal/data/*.go) \
			$(wildcard internal/console/*.go)

.PHONY: all
all: 	$(OUTDIR)/$(APPNAME) \
		$(OUTDIR)/$(APPNAME)_min

.PHONY: clean
clean:
	@rm -rf $(OUTDIR)

$(OUTDIR):
	@mkdir -p $(OUTDIR)

$(OUTDIR)/$(APPNAME): $(OUTDIR) $(SOURCES)
	@echo "building with go..."
	go build -a -gcflags=all="-l -B -wb=false" -ldflags="-s -w" -o $@ main.go
	@$(STRIP) $@
	@$(UPX) $@
	@ls -ahl $@

$(OUTDIR)/$(APPNAME)_min: $(OUTDIR) $(SOURCES)
	@if command -v $(TINYGO) &> /dev/null; then \
  		echo "building with tinygo..."; \
		$(TINYGO) build -gc leaking -scheduler none -opt z -o $@ main.go; \
	else \
	  	echo "building with docker tinygo..."; \
		$(TINYGO_DOCKER) build -gc leaking -scheduler none -opt z -o $@ main.go; \
	fi
	@$(STRIP) $@
	@$(UPX) $@
	@ls -ahl $@

.PHONY: run
run:
	@go run main.go

.PHONY: info
info:
	@echo SOURCES=$(SOURCES)
	@echo TINYGO_DOCKER=$(TINYGO_DOCKER)
	@echo TINYGO=$(TINYGO)
