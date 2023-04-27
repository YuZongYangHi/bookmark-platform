package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type CommonResult struct {
	Success      bool           `json:"success"`
	Data         *ContentResult `json:"data"`
	ErrorCode    int            `json:"errorCode"`
	ErrorMessage string         `json:"errorMessage"`
}

type ContentResult struct {
	List interface{} `json:"list"`
}

type BaseController struct {
	web.Controller
}

func (c *BaseController) JsonResponse(data interface{}) {
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *BaseController) ErrorResponse(code int, msg string) {
	handleResponse := CommonResult{}
	handleResponse.ErrorCode = code
	handleResponse.ErrorMessage = msg
	c.CustomAbort(code, c.ReversedStr(handleResponse))
}

func (c *BaseController) SuccessResponse(data interface{}) {
	handleResponse := CommonResult{}
	handleResponse.Success = true
	handleResponse.Data = &ContentResult{
		List: data,
	}
	c.JsonResponse(handleResponse)
}

func (c *BaseController) BadRequestResponse(msg string) {
	handleResponse := CommonResult{}
	handleResponse.Success = false
	handleResponse.ErrorCode = 400
	handleResponse.ErrorMessage = msg
	handleResponse.Data = &ContentResult{
		List: nil,
	}
	c.JsonResponse(handleResponse)
}

func (c *BaseController) ReversedStr(body interface{}) string {
	res, _ := json.Marshal(body)
	return fmt.Sprintf("%s", res)
}
