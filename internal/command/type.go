package command

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/svang/svangapi/internal/config"
)

//Instance is the main struct for command configs
type Instance struct {
	config    *config.SvangType
	Error     error
	Flags     *pflag.FlagSet
	StartTime time.Time
	ConfigDir string
}
