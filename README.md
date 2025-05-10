# gin-study
使用gin框架进行开发测试学习

# 目录
[1. 使用makefile管理项目](./docs/使用makefile管理项目.md)  
[2. git提交规范限制](./docs/git提交规范限制.md)

# 使用golangci-lint进行静态代码检查
[golangci官方文档](https://golangci-lint.run/)

- 使用 `make install.golangci-lint` 命令单个进行安装， 也可以使用 `make tools.install` 进行安装
- 静态代码检查文件 `.golangci.yaml`。 `output` 指定生成 txt 和 html 格式文件到 `_output/lint` 文件夹下
- 使用 `make go.lint` 进行静态代码检查， 在检查过程中会根据 `formatters` 对所有文件进行格式化，不需要单独进行格式化了。

# 使用make管理docker
使用docker compose来管理需要的docker文件， 因为单个docker-compose文件太大，所以将需要的容器根据名称进行拆分。文件存放在docker文件夹下。
这样不会因为某个容器配置修改导致所有服务启动失败。

容器启动的命令在 scripts/make-rules/docker.mk 中, 可以通过`make help`查看docker相关的命令

# 使用swag生成api文档
[swag文档](https://github.com/swaggo/swag/tree/master/example/celler#parameterType)

- 使用 `make install.swag` 命令进行安装
- 使用 `make go.swag.fmt` 对api相关注释进行格式化
- 使用 `make go.swag.gen` 生成api文档

> 当使用 `make go.lint` 进行静态代码检查时， 可能会删除部分 swag 注释，在做处理时请多注意
