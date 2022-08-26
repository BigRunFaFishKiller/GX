//GX框架的HTTP基础
//使用路由引擎拦截所有请求，并有路由引擎统一处理请求，可以实现自定义路由规则，实现中间件等功能
package GX

import (
	"net/http"
)

type HandlerFunc func(c *Context)

//定义引擎
type Engine struct {
	//将engin作为最顶层的分组
	*RouterGroup
	router *Router
	//保存所有的分组
	groups []*RouterGroup
}

//定义分组路由
type RouterGroup struct {
	//路由前缀
	prefix     string
	middleware []HandlerFunc
	//父分组路由
	parent *RouterGroup
	engine *Engine
}

//创建引擎
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		engine: group.engine,
		prefix: group.prefix + prefix,
		parent: group,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

//给引擎添加路由添加处理器的通用的方法
func (group *RouterGroup) addRouter(method string, pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	group.engine.router.addRouter(method, pattern, handler)
}

//给引擎绑定GET方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRouter("GET", pattern, handler)
}

//给引擎绑定POST方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRouter("POST", pattern, handler)
}

//给引擎绑定DELETE方法
func (group *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	group.addRouter("DELETE", pattern, handler)
}

//给引擎绑定PUT方法
func (group *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	group.addRouter("PUT", pattern, handler)
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
