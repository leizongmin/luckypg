TINYGO ?= docker run --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) tinygo/tinygo:0.23.0 tinygo
OUTDIR ?= $(CURDIR)/out
APPNAME ?= $(shell basename $(CURDIR))

.PHONY: all
all: $(OUTDIR)/$(APPNAME)

.PHONY: clean
clean:
	@rm -rf $(OUTDIR)

$(OUTDIR):
	@mkdir -p $(OUTDIR)

$(OUTDIR)/$(APPNAME): $(OUTDIR) main.go
	@$(TINYGO) build -gc leaking -scheduler none -opt z -o $@ main.go
	@strip $@
	@upx $@
	@ls -ahl $@

.PHONY: run
run:
	@go run main.go
