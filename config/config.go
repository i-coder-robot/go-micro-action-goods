package config

import (
	"github.com/i-coder-robot/go-micro-action-core/config"
	"github.com/i-coder-robot/go-micro-action-core/env"
)

// Conf 配置
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.srv.") + "product",
	Version: "latest",
}