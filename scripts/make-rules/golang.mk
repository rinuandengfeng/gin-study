
## go.lint: 静态代码检查
.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@echo "===========> 运行 golangci 进行静态代码检查"
	golangci-lint run -c $(ROOT_DIR)/.golangci.yaml $(ROOT_DIR)/...