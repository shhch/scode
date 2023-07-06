package scode

type SrcCode struct {
	code    int
	message string
}

// Use NewSrcCode to define service code.
func NewSrcCode(code int, msg string) *SrcCode {
	return &SrcCode{
		code:    code,
		message: msg,
	}
}

func (sc *SrcCode) Code() int {
	return sc.code
}

func (sc *SrcCode) Error() string {
	return sc.message
}
