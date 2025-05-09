
#postgresql
DOCKERS := mariadb redis

## docker.all: 启动所有容器
.PHONY: docker.all
docker.all: $(addprefix docker.,$(DOCKERS))

## docker.all.del: 删除所有容器
.PHONY: docker.all.del
docker.all.del: $(addsuffix .del,$(addprefix docker.,$(DOCKERS)))

## docker.[容器]: 启动docker文件夹下同名yaml文件
.PHONY: docker.%
docker.%:
	@echo "启动 $(*) 容器"
	@$(DOCKER) compose -f $(DOCKER_DIR)/$(*).yaml up -d


## docker.[容器].del: 停止容器并删除
.PHONY: docker.%.del
docker.%.del:
	@echo "正在停止 $(*) 容器"
	@$(DOCKER) compose -f $(DOCKER_DIR)/$(*).yaml down

## docker.[容器].delfile: 停止容器并删除data文件夹下的数据挂载
.PHONY: docker.%.delfile
docker.%.delfile:
	@$(MAKE) docker.$(*).del
	@rm -rf $(DOCKER_DIR)/data/$(*)

## docker.[容器].restart: 重启容器
.PHONY: docker.%.restart
docker.%.restart:
	@$(MAKE) docker.$(*).del
	@$(MAKE) docker.$(*)
