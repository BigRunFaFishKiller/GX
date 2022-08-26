//GX框架的contex上下文
//原生http框架提供的接口颗粒度太细，每次对请求响应时需要进行过于繁琐的封装
//对于从请求中获取的信息，以及中间处理的所产生的信息也需要一个合适的容器承载
//context上下文伴随请求的出现而产生，请求结束而销毁，针对指定的场景进行封装
package GX

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//http请求接口
	Writer http.ResponseWriter
	Req    *http.Request
	//请求的详细信息
	Path       string
	Method     string
	StatusCode int
	//请求路径中的参数
	Params map[string]string
}

//构造Context
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

//获取动态路由中的参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

//获取POST请求表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//获取请求路径中的参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//设置响应码
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//响应String
func (c *Context) String(code int, format string, values ...string) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values)))
}

//响应JSON
func (c *Context) JSON(code int, obj interface{}) {
	encoder := json.NewEncoder(c.Writer)
	//解析JSON失败响应错误
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		c.SetStatus(code)
		c.SetHeader("Content-Type", "application/json")
	}
}

//响应HTML
func (c *Context) HTML(code int, html string) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(html))
}
