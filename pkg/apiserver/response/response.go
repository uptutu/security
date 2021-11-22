/*
Copyright 2021 The tKeel Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package response

import (
	"encoding/json"
	"net/http"

	"github.com/tkeel-io/security/pkg/errcode"

	"github.com/emicklei/go-restful"
)

type responseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SrvErrWithRest(resp *restful.Response, err *errcode.SrvError, data interface{}) {
	resp.WriteEntity(responseData{
		Code: err.Code(),
		Msg:  err.Msg(),
		Data: data,
	})
}

func DefineWithRest(resp *restful.Response, code int, msg string, data interface{}) {
	resp.WriteEntity(responseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func SrvErrWithWriter(w http.ResponseWriter, err *errcode.SrvError, data interface{}) {
	resp := responseData{
		Code: err.Code(),
		Msg:  err.Msg(),
		Data: data,
	}

	respBytes, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func DefineWithWriter(w http.ResponseWriter, code int, msg string, data interface{}) {
	resp := responseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	respBytes, _ := json.Marshal(resp)
	w.Header().Set("ContentType", "application/json")
	w.Write(respBytes)
}
