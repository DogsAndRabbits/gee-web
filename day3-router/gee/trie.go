package gee

import(
	"fmt"
	"strings"
)

type node struct{
	pattern string
	part string
	children []*node
	isWild bool
}

func (n *node)matchChild(part string)*node{
	for _,child:=range n.children{
		if child.part==part || child.isWild{
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node )insert(pattern string,parts []string,height int){
	nodes:=make([]*node,0)
	for _,child:=range n.children{
		if child.part == part || child.isWild{
			nodes = append(nodes,child)
		}
	}
	return nodes
}

