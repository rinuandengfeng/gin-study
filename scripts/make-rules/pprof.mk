PPROF_DIR = $(OUTPUT_DIR)/pprof

# 统一定义采集文件生成信息
PPROF_FLAGS = -gcflags "all=-N -l"
# 生成cpu.profile
PPROF_FLAGS += -cpuprofile $(PPROF_DIR)/$(Func)cpu.profile
# 生成mem.profile
PPROF_FLAGS += -memprofile $(PPROF_DIR)/$(Func)mem.profile
# 生成block.profile
PPROF_FLAGS += -blockprofile $(PPROF_DIR)/$(Func)block.profile
PPROF_FLAGS += -o $(PPROF_DIR)/$(Func)$(SUF)


## :
## pprof.test: 生成单个单元测试的性能数据文件，Func 要进行单元测试的方法，Package 要进行测试的方法包位置。
## : Func 为空时会执行 Package 中的所有测试方法
## : 例：make pprof.test Func=TestSendCode Package=gin-study/internal/controller/user
.PHONY: pprof.test
pprof.test:
	@mkdir -p $(PPROF_DIR)
	$(GO) test -run $(Func) -v $(Package) $(PPROF_FLAGS)

## pprof.bench: 生成基准测试的性能测试数据 参数信息与pprof.test一致
## : make pprof.test Func=BenchmarkSendCode Package=gin-study/internal/controller/user
.PHONY: pprof.bench
pprof.bench:
	@mkdir -p $(PPROF_DIR)
	# -benchmem 显示与内存分配相关的详细信息
	# -benchtime 设定每个基准测试用例的运行时间
	$(GO) test -bench $(Func) -v $(Package) -benchmem -benchtime=10s $(PPROF_FLAGS)

## pprof.mem: 使用 8001 端口查看profile文件分析, 需要使用到 Func 参数, Func要和 pprof.test 中使用的Func一致
## : 例：make pprof.cpu Func=TestSendCode
## : 例：make pprof.mem Func=TestSendCode
## : 例：make pprof.block Func=TestSendCode
.PHONY: pprof.%
pprof.%:
	@$(GO) tool pprof -http=":8001" $(PPROF_DIR)/$(Func)$(*).profile

## pprof.trace: 生成trace文件
.PHONY: pprof.trace
pprof.trace:
	@mkdir -p $(PPROF_DIR)
	@curl "http://127.0.0.1:8000/debug/pprof/trace?second=30" > $(PPROF_DIR)/trace.out


.PHONY: pprof.trace.show
pprof.trace.show:
	@$(GO) tool trace $(PPROF_DIR)/trace.out