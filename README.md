# scode
http server code


## scode example usage:
```golang
package main

import (
    "github.com/shhch/scode"
    "errors"
)

var (
    SUCCESS       = scode.NewSrcCode(0, "success")
    OTHER_ERROR   = scode.NewSrcCode(40001, "other error")
    SYSTEM_ERROR  = scode.NewSrcCode(-1, "system error")
)

var testErr = errors.New("test error")

func main() {
    scode.SetCaller(true)
    err := biz()
    if err != nil {
        code, msg, trace := parseCode(err)
        println(code)
        println(msg)
        println(trace)
    }
}

func biz() error {
    err := test()
    if err != nil {
        return OTHER_ERROR.C().DescE(err, "test error")
    }
    return nil
}

func test() error {
    return testErr
}

func parseCode(err error) (code int, message, traceMsg string) {
    if err == nil {
        return
    }
    
    switch err.(type) {
    case *scode.SCode:
        v := err.(*scode.SCode)
        code = v.Code()
        message = v.Error()
        traceMsg = v.SourceErrMsg()
    case *scode.SrcCode:
        v := err.(*scode.SrcCode)
        code = v.Code()
        message = v.Error()
        traceMsg = message
    case error:
        code = SYSTEM_ERROR.Code()
        message = SYSTEM_ERROR.Error()
        traceMsg = err.Error()
    }
    return
}

```
