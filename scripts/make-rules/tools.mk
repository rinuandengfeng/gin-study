# 安装项目依赖工具
TOOLS ?= golangci-lint swag mockgen go-junit-report

## :
## tools.install: 安装所有工具
.PHONY: tools.install
tools.install: $(addprefix tools.install.,$(TOOLS))

# 使用 % 通配符对安装工具进行统一输出和控制
.PHONY: tools.install.%
tools.install.%:
	@echo "====================> 安装 $*"
	@$(MAKE) install.$*

# 校验工具是否安装
.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi


# 安装静态代码检查工具
.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.2

# 安装swagger生成器
.PHONY: install.swag
install.swag:
	$(GO) install github.com/swaggo/swag/cmd/swag@latest

# 安装mock工具生成单元测试依赖代码
.PHONY: install.mockgen
install.mockgen:
	@$(GO) install go.uber.org/mock/mockgen@latest

# 安装 将 go test 输出转换为xml工具
.PHONY: install.go-junit-report
install.go-junit-report:
	@$(GO) install github.com/jstemmer/go-junit-report@latest