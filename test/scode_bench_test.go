package main

import (
	"errors"
	"github.com/shhch/scode"
	"github.com/shhch/scode/example"
	"testing"
)

func BenchmarkNoCaller(b *testing.B) {
	scode.SetCaller(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		example.SYSTEM_ERROR.C().StrE("test error").SourceErrMsg()
	}
}

func BenchmarkCaller(b *testing.B) {
	scode.SetCaller(true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		example.SYSTEM_ERROR.C().StrE("test error").SourceErrMsg()
	}
}

func BenchmarkError(b *testing.B) {
	scode.SetCaller(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		example.SYSTEM_ERROR.C().E(errors.New("test error")).SourceErrMsg()
	}
}

func BenchmarkErrorWithStr(b *testing.B) {
	scode.SetCaller(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		example.SYSTEM_ERROR.C().DescE(errors.New("test error"), "desc").SourceErrMsg()
	}
}
