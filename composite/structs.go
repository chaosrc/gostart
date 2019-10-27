package main

import (
	"fmt"
)

type Employee struct {
	ID      int
	Name    string
	Address string
	Salary  int
	time    int
}

var foo Employee

func field() {
	foo.Salary = 5000
	name := &foo.Name

	*name = *name + "lastName"

	p := &foo
	p.Address = "shanghai"

	fmt.Println(foo)
}

// Tree for tree
type Tree struct {
	value       int
	left, right *Tree
}

func sort(values []int) {
	var root *Tree

	for _, value := range values {
		root = add(root, value)
	}

	toArray(values[:0], root)
}

func add(tree *Tree, value int) *Tree {
	if tree == nil {
		tree = &Tree{value: value}
		return tree
	}

	if value <= tree.value {
		tree.right = add(tree.right, value)
	}
	if value > tree.value {
		tree.left = add(tree.left, value)
	}
	return tree
}

func toArray(list []int, tree *Tree) []int {
	if tree == nil {
		return list
	}
	if tree.right != nil {
		list = toArray(list, tree.right)
	}
	list = append(list, tree.value)
	if tree.left != nil {
		list = toArray(list, tree.left)
	}
	return list
}
