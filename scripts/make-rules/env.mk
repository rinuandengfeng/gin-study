SHELL := /bin/bash
GO := go
GOOS = $(shell $(GO) env GOOS)
FIND := find
GIT := git

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../.. && pwd -P))

ifeq ($(GOOS),windows)
	# 如果 ROOT_DIR 位置不正确可以在此处直接指定文件位置
	ROOT_DIR = E:/code/gin-study
	FIND="D:\Program Files\Git\usr\bin\find"
endif

OUTPUT_DIR := $(ROOT_DIR)/_output
