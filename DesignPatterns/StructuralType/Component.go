package StructuralType

import "fmt"

type Component interface {
	// 采购设备或者添加子部门
	Add(component Component)
	Remove(component Component)
	// 查询该节点下所有设备和部门
	Display(depth int)
}

type Leaf struct {
	name string
}

func (l *Leaf) Add(component Component) {
	panic("叶子节点不能挂载设备")
}

func (l *Leaf) Remove(component Component) {
	panic("叶子节点不能移除设备")
}

func (l *Leaf) Display(depth int) {
	// 输出树形结构的叶子结点，这里直接输出设备名
	for i:=0; i<depth; i++ {
		fmt.Print("*")
	}
	fmt.Println(l.Name())
}

func (l *Leaf) Name() string {
	return l.name
}

func (l *Leaf) SetName(name string) {
	l.name = name
}

// 复合构件
type Composite struct {
	name string
	arr []Component
}

func (c *Composite) Add(component Component) {
	c.arr = append(c.arr,component)
}

func (c *Composite) Remove(component Component) {
	for i,v := range c.arr {
		if v == component {
			// 删除第i个元素,因为interface类型在golang中
			// 以地址的方式传递，所以可以直接比较进行删除
			// golang中只要记得byte,int,bool,string，数组，结构体，默认传值，其他的默认传地址即可
			c.arr = append(c.arr[:i],c.arr[i+1:]...)
		}
	}
}

func (c *Composite) Display(depth int) {
	// 输出树形结构
	for i:=0; i<depth; i++ {
		fmt.Print("*")
	}
	fmt.Println(c.Name())
	// 递归显示
	for _,com := range c.arr {
		com.Display(depth+1)
	}
}

func (c *Composite) Name() string {
	return c.name
}

func (c *Composite) SetName(name string) {
	c.name = name
}
