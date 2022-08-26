//GX框架的HTTP基础
//使用路由引擎拦截所有请求，并有路由引擎统一处理请求，可以实现自定义路由规则，实现中间件等功能
package GX

import (
	"net/http"
)

type HandlerFunc func(c *Context)

//定义引擎
type Engine struct {
	router *Router
}

//创建引擎
func New() *Engine {
	return &Engine{router: newRouter()}
}

//给引擎添加路由添加处理器的通用的方法
func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

//给引擎绑定GET方法
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

//给引擎绑定POST方法
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

//给引擎绑定DELETE方法
func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.addRouter("DELETE", pattern, handler)
}

//给引擎绑定PUT方法
func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.addRouter("PUT", pattern, handler)
}

//引擎实现handle接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	e.router.Handle(context)
}

//启动Web服务器
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
