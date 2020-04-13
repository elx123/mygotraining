package main

import(
"fmt"
"container/list"
)

type node struct{
	data int 
	left *node
	right *node
}

var root *node

func qianxu(root *node){
	if(root == nil){
		return
	}
	fmt.Println(root.data)
	qianxu(root.left)
	qianxu(root.right)
}

func zhongxu(root *node){
	if(root == nil){
		return
	}
	qianxu(root.left)
	fmt.Println(root.data)
	qianxu(root.right)
}

func houxu(root *node){
	if(root == nil){
		return
	}
	qianxu(root.left)
	qianxu(root.right)
	fmt.Println(root.data)
}

func generate()*node{
	a:= node{data:10,}
	b:= node{data:11,}
	c:= node{data:12,}
	d:= node{data:13,}
	e:= node{data:14,}
	a.left = &b
	a.right = &c
	b.left = &d
	c.right = &e
	return &a
}

func nodigui(root *node){
	stack := list.New()

	for stack.Len()>0 || root != nil{
		if root != nil{
			stack.PushFront(root)
			root = root.left
		}else{
			root = stack.Front().Value.(*node)
			fmt.Println(root.data)
			root = root.right

			stack.Remove(stack.Front())
		}
	}
}

func main(){
	root := generate()
	zhongxu(root)
	fmt.Println("------------------------------------------------------------------")
	nodigui(root)

}