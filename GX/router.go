package GX

import "net/http"

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//解析请求路径，寻找对应的处理器方法，并执行
func (r *Router) Handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s", c.Path)
	}
}
