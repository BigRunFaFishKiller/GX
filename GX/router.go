package GX

import (
	"net/http"
	"strings"
)

type Router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func newRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

//以"/"为分割符，将路径分割
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {

	parts := parsePattern(pattern)
	key := method + "-" + pattern

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

//查询所给的路径是否存在，并返回对应路劲构成的前缀树和每个节点构成的map
func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			//获取请求路径中的参数
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

//解析请求路径，寻找对应的处理器方法，并执行
func (r *Router) Handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		//将处理器方法放入HandlerFunc切片
		c.Handlers = append(c.Handlers, r.handlers[key])
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s", c.Path)
	}
	//开始时执行中间件，中间件的最后一个方法为处理器方法
	c.Next()
}
