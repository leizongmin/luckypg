GO ?= go
TINYGO ?= tinygo
STRIP ?= strip

OUTDIR ?= $(CURDIR)/out
APPNAME ?= $(shell basename $(CURDIR))
SOURCES ?= 	$(wildcard *.go) \
			$(wildcard internal/nostd/*.go) \
			$(wildcard internal/data/*.go) \
			$(wildcard internal/console/*.go)

.PHONY: all
all: tiny

.PHONY: clean
clean:
	@rm -rf $(OUTDIR)

$(OUTDIR):
	@mkdir -p $(OUTDIR)

$(OUTDIR)/$(APPNAME)_normal: $(OUTDIR) $(SOURCES)
	@echo "building with $(GO)..."
	@$(GO) build -a -gcflags=all="-l -B -wb=false" -ldflags="-s -w" -o $@ main.go
	@$(STRIP) $@
	@ls -ahl $@

$(OUTDIR)/$(APPNAME): $(OUTDIR) $(SOURCES)
	@echo "building with $(TINYGO)..."
	$(TINYGO) build -gc leaking -scheduler none -opt z -o $@ main.go
	@$(STRIP) $@
	@ls -ahl $@

.PHONY: run
run:
	@$(GO) run main.go

.PHONY: tiny
tiny: $(OUTDIR)/$(APPNAME)

.PHONY: normal
normal: $(OUTDIR)/$(APPNAME)_normal

.PHONY: info
info:
	@echo SOURCES=$(SOURCES)
	@echo TINYGO=$(TINYGO)
	@echo GO=$(GO)
	@echo STRIP=$(STRIP)
	@echo OUTDIR=$(OUTDIR)
	@echo OUTDIR=$(OUTDIR)
	@echo APPNAME=$(APPNAME)
