package utils

import (
	"os"
	"os/signal"
	"syscall"
)

type (
	// IShutDown 优雅关闭服务的接口,.
	IShutDown interface {
		// Close 注册关闭服务
		Close()
		// Register 注册需要进行关闭的服务信息
		Register(fs ...func())
	}

	shutDown struct {
		signals chan os.Signal
		fs      []func()
	}
)

// NewShutDown 创建服务, 默认监听SIGINT和SIGTERM.
func NewShutDown(signals ...syscall.Signal) IShutDown {
	s := &shutDown{signals: make(chan os.Signal, 1)}
	for _, sig := range signals {
		signal.Notify(s.signals, sig)
	}

	return s
}

// Register 注册需要进行关闭的服务信息.
func (s *shutDown) Register(fs ...func()) {
	s.fs = append(s.fs, fs...)
}

// Close 注册关闭服务.
func (s *shutDown) Close() {
	<-s.signals
	signal.Stop(s.signals)

	for _, f := range s.fs {
		f()
	}
}
