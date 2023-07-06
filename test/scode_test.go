package main

import (
	"errors"
	"fmt"
	"github.com/shhch/scode"
	"github.com/shhch/scode/example"
	"testing"
)

var (
	errmsg  = "test error"
	testErr = errors.New(errmsg)
)

func TestScodeNilErr(t *testing.T) {
	scode.SetCaller(true)
	fmt.Println(example.SYSTEM_ERROR.C().SourceErrMsg())
}

func TestScodeErr(t *testing.T) {
	scode.SetCaller(true)
	fmt.Println(example.SYSTEM_ERROR.C().E(testErr).SourceErrMsg())
}

func TestScodeStrErr(t *testing.T) {
	scode.SetCaller(true)
	fmt.Println(example.SYSTEM_ERROR.C().StrE(errmsg).SourceErrMsg())
}

func TestScodeDescErr(t *testing.T) {
	scode.SetCaller(true)
	fmt.Println(example.SYSTEM_ERROR.C().DescE(testErr, "desc").SourceErrMsg())
}

func TestScodeDescStrErr(t *testing.T) {
	scode.SetCaller(true)
	fmt.Println(example.SYSTEM_ERROR.C().DescStrE(errmsg, "desc").SourceErrMsg())
}

func TestScodeCaller(t *testing.T) {
	scode.SetCaller(true)
	fmt.Println(example.SYSTEM_ERROR.C().DescE(testErr, "desc").SourceErrMsg())
}
