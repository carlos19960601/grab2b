.PHONY: binaries

ICON="🔞"

# 项目的二进制命令
COMMANDS=grab2b show2b

BINARIES=$(addprefix bin/,$(COMMANDS))

all: binaries

FORCE:
define BUILD_BINARY
@echo "$(ICON) $@"
@go build -o $@ ./$<
endef

binaries: $(BINARIES) ## build binaries
	@echo "$(ICON) $@"

bin/%: cmd/% FORCE
	$(call BUILD_BINARY)