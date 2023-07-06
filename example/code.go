package example

import "github.com/shhch/scode"

var (
	SUCCESS       = scode.NewSrcCode(0, "success")
	SYSTEM_ERROR  = scode.NewSrcCode(-1, "system error")
	PARAM_ERROR   = scode.NewSrcCode(4001, "param error")
	OPERATE_ERROR = scode.NewSrcCode(5001, "operate error")
)
