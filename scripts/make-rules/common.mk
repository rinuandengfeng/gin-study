
# 复制githooks文件夹下的钩子到 .git/hooks文件夹下
COPY_GITHOOK:=$(shell cp -f scripts/githooks/* .git/hooks/)
