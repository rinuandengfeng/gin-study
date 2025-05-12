
## :
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

# 正则表达式排除不进行覆盖率测试的包
EXCLUDE_TESTS = 'gin-study/(api|internal/(cache/mock|store/mock))'

## go.test: 运行项目中所有的单元测试，排除 mock 和 EXCLUDE_TESTS指定文件夹
.PHONY: go.test
go.test: tools.verify.go-junit-report
	@echo "===========> 单元测试中..."
	@mkdir -p $(OUTPUT_DIR)
	@set -o pipefail;$(GO) test -race -cover -coverprofile=$(OUTPUT_DIR)/coverage.out \
		-timeout=10m -shuffle=on -short -v `go list ./... | grep -vE $(EXCLUDE_TESTS)` 2>&1 | \
		tee >(go-junit-report --set-exit-code >$(OUTPUT_DIR)/report.xml)
	@sed -i '/mock_.*.go/d' $(OUTPUT_DIR)/coverage.out # 从 coverage 中删除mock_.*.go 文件
	@$(GO) tool cover -html=$(OUTPUT_DIR)/coverage.out -o $(OUTPUT_DIR)/coverage.html


## go.test.cover: 单元测试覆盖率判断, 覆盖率默认60, 使用COVERAGE进行修改。
## : 例：make go.test.conver COVERAGE=50
.PHONY: go.test.cover
go.test.cover: go.test
	@$(GO) tool cover -func=$(OUTPUT_DIR)/coverage.out  | awk -v target=$(COVERAGE) -f $(ROOT_DIR)/scripts/coverage.awk