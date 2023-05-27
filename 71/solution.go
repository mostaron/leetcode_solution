package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))

}

type Node struct {
	value string
	pre   *Node
}

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	var tail *Node
	for _, dir := range dirs {
		if dir == "." || dir == "" {
			continue
		}
		if dir == ".." {
			if tail != nil {
				tail = tail.pre
			}
			continue
		}
		if tail == nil {
			tail = &Node{value: dir, pre: nil}
		} else {
			tail = &Node{value: dir, pre: tail}
		}
	}
	result := ""
	for tail != nil {
		result = "/" + tail.value + result
		tail = tail.pre
	}
	if result == "" {
		return "/"
	} else {
		return result
	}
}
