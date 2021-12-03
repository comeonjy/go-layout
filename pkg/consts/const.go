package consts

import (
	"github.com/comeonjy/go-kit/pkg/xenv"
)

var EnvMap = map[string]string{
	xenv.AppName:     "go-layout",
	xenv.AppVersion:  "v1.0",
	xenv.TraceName:   "trace_id",
	xenv.ApolloAppID: "go-layout",
	xenv.ApolloUrl:   "http://apollo.dev.jiangyang.me",
	"my_const":       "my_const_value",
}
