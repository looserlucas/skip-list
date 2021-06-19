package main

import (
	"fmt"
	"math/rand"
)

const headValue = "Head-Value"
const maxHeight = 6

type SkipList struct {
	Head *SkipListNode
}

type SkipListNode struct {
	Value     string
	Height    int
	Successor []*SkipListNode
}

func Create() (list *SkipList, err error) {
	// create head node
	head := &SkipListNode{
		Value:     headValue,
		Height:    maxHeight,
		Successor: make([]*SkipListNode, maxHeight),
	}

	list = &SkipList{
		Head: head,
	}
	return list, nil
}

func (n *SkipList) AddNode(value string) (ok bool, err error) {
	//check if node exists
	node, ok, err := n.Travel(value, false)
	//if exist skip
	if ok && (node != nil) {
		return false, fmt.Errorf("This node has existed in the list")
	}
	//if err skip
	if !ok && (err != nil) && (node == nil) {
		return false, err
	}

	//else create new node and add to skiplist
	newNode := &SkipListNode{
		Value:     value,
		Height:    rand.Intn(5),
		Successor: make([]*SkipListNode, maxHeight),
	}

	//fix pointers
	for i := 0; i < newNode.Height; i++ {
		newNode.Successor[i] = node[i].Successor[i]
		node[i].Successor[i] = newNode
	}
	return
}

func (n *SkipList) DelNode(value string) (ok bool, err error) {
	//check if node exists
	node, ok, err := n.Travel(value, false)
	if !ok {
		return false, fmt.Errorf("There is no node to delete")
	}
	if ok && (err != nil) && (node == nil) {
		return false, err
	}

	//delete node
	//fix pointer
	for i := 0; i < node[0].Height; i++ {
		node[i].Successor[i] = node[i].Successor[i].Successor[i]
	}

	return
}

func (n *SkipList) Find(value string) (node *SkipListNode, ok bool, err error) {
	//find the node with exact value
	//else return ok=false, err=nil, node = node with closest value which < value
	nodes, ok, err := n.Travel(value, true)
	if !ok {
		return nil, false, fmt.Errorf("There is no node to delete")
	}
	if ok && (err != nil) && (node == nil) {
		return nil, false, err
	}
	node = nodes[0]

	return
}

// Travel: Return the neighbors with the level from (0 to maxHeight) of the value node
// This function have 2 mode:
// find mode to find the exact node.
// travel mode to find the neighbors with the level from (0 to maxHeight) of the value node
func (n *SkipList) Travel(value string, find bool) (node []*SkipListNode, ok bool, err error) {
	node = make([]*SkipListNode, maxHeight)
	// start with head, and max level
	thisNode := n.Head
	level := maxHeight

	for {
		// break condition
		if (level == 0) && (thisNode.Successor[0].Value > value) {
			node[0] = thisNode
			break
		}

		// if next node have the value < target value, keep going on finding on this level
		if thisNode.Successor[level].Value < value {
			thisNode = thisNode.Successor[level]
		} else {
			//if we have found the node with the target value
			if thisNode.Successor[level].Value == value {
				// if it s find function, return this node
				if find {
					var realNode []*SkipListNode
					realNode = append(realNode, thisNode.Successor[level])
					return realNode, true, nil
				} else {
					// if it is not find function, keep going on finding the neighbor
					node[level] = thisNode
					level -= 1
					ok = true
				}
			} else {
				// if the next node have the value > target value, move down 1 level
				node[level] = thisNode
				if level > 0 {
					level -= 1
				}
			}
		}

	}
	return node, false, nil
}
func main() {
}
