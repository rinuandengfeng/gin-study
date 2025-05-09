SHELL := /bin/bash
GO := go
GOOS = $(shell $(GO) env GOOS)
FIND := find
GIT := git

ifeq ($(GOOS),windows)
	FIND="D:\Program Files\Git\usr\bin\find"
endif