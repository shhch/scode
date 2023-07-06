package main

import (
	"encoding/json"
	"errors"
	"github.com/shhch/scode"
	"github.com/shhch/scode/example"
	"net/http"
)

func testErr() error {
	return errors.New("test error")
}

type RespBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func writeResp(w http.ResponseWriter, data interface{}, sc error) {
	if data == nil {
		data = struct{}{}
	}
	resp := &RespBody{
		Data: data,
	}

	traceMsg := ""
	resp.Code, resp.Message, traceMsg = parseCode(sc)
	if resp.Code != 0 {
		print(traceMsg)
	}

	body, _ := json.Marshal(resp)
	println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
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
		code = example.SYSTEM_ERROR.Code()
		message = example.SYSTEM_ERROR.Error()
		traceMsg = err.Error()
	}
	return
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	if vals.Get("test") == "" {
		writeResp(w, nil, example.PARAM_ERROR.C())
		return
	} else if vals.Get("test") != "test" {
		err := testErr()
		if err != nil {
			writeResp(w, nil, example.OPERATE_ERROR.C().DescE(err, "testErr"))
			return
		}
	}

	writeResp(w, nil, example.SUCCESS)
	return
}

func main() {
	scode.SetCaller(true)
	http.HandleFunc("/test", handleTest)
	http.ListenAndServe(":8583", nil)
}
