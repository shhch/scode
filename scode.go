package scode

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// SCode encapsulates SrcCode, allowing additional business
// information and stack information to be appended.
// It must be used at the lowest level of error occurrence
// without duplication.
// Use SourceErrMsg to obtain a detailed description
// and free it during final processing.
type SCode struct {
	srcc      *SrcCode
	sourceErr error
	caller    string
	bizDesc   string
}

var (
	withCaller = false
	srcp       = sync.Pool{
		New: func() any {
			return new(SCode)
		},
	}
)

func SetCaller(open bool) {
	withCaller = open
}

// Use the func C() to new an SCode structure.
func (c *SrcCode) C() *SCode {
	sc := srcp.Get().(*SCode)
	sc.srcc = c
	sc.sourceErr = nil
	return sc
}

func (sc *SCode) free() {
	sc.srcc = nil
	sc.sourceErr = nil
	sc.bizDesc = ""
	sc.caller = ""
	srcp.Put(sc)
}

func (sc *SCode) SourceErrMsg(withoutfree ...bool) string {
	if len(withoutfree) == 0 || !withoutfree[0] {
		defer sc.free()
	}

	b := strings.Builder{}
	if sc.caller != "" {
		b.WriteString(sc.caller + " ")
	}
	if sc.srcc != nil {
		b.WriteString(sc.srcc.Error() + "|")
	}
	if sc.bizDesc != "" {
		b.WriteString(sc.bizDesc + ":")
	}
	if sc.sourceErr != nil {
		b.WriteString(sc.sourceErr.Error())
	}
	return b.String()
}

func (sc *SCode) Code() int {
	return sc.srcc.Code()
}

func (sc *SCode) Error() string {
	return sc.srcc.Error()
}

func (sc *SCode) addcaller() {
	_, fpath, line, _ := runtime.Caller(2)
	pathArr := strings.Split(fpath, "/")
	sc.caller = fmt.Sprintf("%s:%d", pathArr[len(pathArr)-1], line)
}

func (sc *SCode) E(err error) *SCode {
	if withCaller {
		sc.addcaller()
	}
	sc.sourceErr = err
	return sc
}

func (sc *SCode) DescE(err error, desc string) *SCode {
	if withCaller {
		sc.addcaller()
	}
	sc.sourceErr = err
	sc.bizDesc = desc
	return sc
}

func (sc *SCode) StrE(errMsg string) *SCode {
	if withCaller {
		sc.addcaller()
	}
	sc.sourceErr = errors.New(errMsg)
	return sc
}

func (sc *SCode) DescStrE(errMsg string, desc string) *SCode {
	if withCaller {
		sc.addcaller()
	}
	sc.sourceErr = errors.New(errMsg)
	sc.bizDesc = desc
	return sc
}

func (sc *SCode) IsSrcCode(code *SrcCode) bool {
	return errors.Is(sc.srcc, code)
}
