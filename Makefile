.DEFAULT_ALL = all

include ./scripts/make-rules/env.mk
include ./scripts/make-rules/common.mk

## help: 帮助命令
.PHONY: help
help:
	@echo "操作系统：$(GOOS)"
	@printf "\n使用: make <TARGETS> \n\n命令:\n"
	@sed -n 's/^##//p' Makefile | column -t -s ':' | sed -e 's/^/ /'
	@$(FIND) ./scripts/make-rules -name "*.mk" | xargs sed -n 's/^##//p' | column -t -s ':' | sed -e 's/^/ /'

