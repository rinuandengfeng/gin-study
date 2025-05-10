package options

import "time"

type ServerOption struct {
	Name         string        `json:"name"          mapstructure:"name"`
	Address      string        `json:"address"       mapstructure:"address"`
	RunMode      string        `json:"run_mode"      mapstructure:"run_mode"`
	ReadTimeout  time.Duration `json:"read_timeout"  mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
}
