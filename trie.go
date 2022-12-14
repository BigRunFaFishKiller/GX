//使用前缀树完成对匹配动态路由
package GX

import "strings"

//trie的节点
type node struct {
	//带匹配的路由，例如/p/:lang
	pattern string
	//路由中的一部分，例如:lang，一般用于表示当前层的节点
	part string
	//子节点
	children []*node
	//是否精确匹配，part含有*或：即精确匹配
	isWild bool
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//递归查找每一层节点，如果没有节点，就会新建一个
//注意：对于路径：/p/:lang/doc，只有在第三层节点即/doc节点才会给pattern设置/p/:lang/doc，/p和/:lang节点的值会被设置为空
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

//根据part即从请求路径中分割出的节点切片中，查找最末端的pattern不为空节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
