
## go.lint: 静态代码检查
.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@echo "===========> 运行 golangci 进行静态代码检查"
	golangci-lint run -c $(ROOT_DIR)/.golangci.yaml $(ROOT_DIR)/...

## go.tidy: 执行 go mod tidy
.PHONY: go.tidy
go.tidy:
	@$(GO) mod tidy

## go.swag.fmt: 格式化swag注释信息
.PHONY: go.swag.fmt
go.swag.fmt: tools.verify.swag
	swag fmt -d ./ --exclude ./pkg,./scripts

## go.swag.gen: 生成swag文档
.PHONY: go.swag.gen
go.swag.gen: tools.verify.swag
	swag init -o $(ROOT_DIR)/api