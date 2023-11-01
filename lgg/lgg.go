package lgg

import (
	"heya/config"
	"heya/lgg/internal"
)

var LogLevel = internal.FromString(config.Config.LogLevelStr)

var LGG = internal.NewLogger(LogLevel, nil, nil)
var New = internal.NewLogger
var Errorf = LGG.Errorf
var Error = LGG.Error
var Infof = LGG.Infof
var Info = LGG.Info
var Debugf = LGG.Debugf
var Debug = LGG.Debug

var Panic = LGG.Panic
var Panicf = LGG.Panicf
