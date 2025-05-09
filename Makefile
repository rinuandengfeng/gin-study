.DEFAULT_ALL = all

include ./scripts/make-rules/env.mk
include ./scripts/make-rules/common.mk
include ./scripts/make-rules/golang.mk
include ./scripts/make-rules/tools.mk

## all: 执行静态代码检查
.PHONY: all
all: go.tidy go.lint

## clean: 清除生成到 OUTPUT_DIR 文件夹下的文件
.PHONY: clean
clean:
	@rm -rf $(OUTPUT_DIR)

## help: 帮助命令
.PHONY: help
help:
	@printf "\n使用: make <TARGETS> \n\n命令:\n"
	@sed -n 's/^##//p' Makefile | column -t -s ':' | sed -e 's/^/ /'
	@$(FIND) ./scripts/make-rules -name "*.mk" | xargs sed -n 's/^##//p' | column -t -s ':' | sed -e 's/^/ /'

