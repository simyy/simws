package tair

import (
    "fmt"
)

type Tair struct {
    root Node
    count int
}

func (t *Tair) New() {
    t.root = Node.New(nil)
    t.count = 0
}

func (t *Tair) Find(values []byte) Node {
    if t.count == 0 {
        return nil
    }

    return t.root.find(values, 0)
}

func (t *Tair) Add(values []byte) {
    t.root.add(values []byte, 0)
    t.count += 1
}

func (t *Tair) Del(values []byte) {
    if t.Find(values) == true {
        t.root.del(values []byte, 0)
        t.count -= 1
    }
}

type Node struct {
    value byte
    nextNodes []*Node
    nums int
}

func (t *Node) new(value byte) {
    return Node{
        value: value,
        nextNodes: nil,
        nums: 0 
    }
}

func (t *Node) add(values []byte, place int) {
    if place == len(values) {
        return
    }

    if t.nextNodes == nil {
        t.nextNodes = make([]Node, 26, 5)
    }

    for _, node := t.nextNodes {
        if node.values == values[place] {
            node.nums += 1
            node.add(values []byte, place + 1)
            return
        }
    }
    node = Node.new(values[place])
    t.nextNodes = append(t.nextNodes, node)
    node.add(values, place + 1)
}

func (t *Node) del(values []byte, place int) {
    if t.nums == 0 || place == len(values) {
        return
    }

    for _, node := t.nextNodes {
        if node.values == values[place] {
            t.nums -= 1
            node.del(values, place + 1)
            if node.nums < 0 {
                t.remove(node) // replace slice, use *list.List 
            }
        }
        break
    }
}

func (t *Node) find(values []byte, place int) bool {
    if t == nil || t.nums == 0 {
        return false
    }

    if place >= len(values) {
        return false
    }

    for _, node := t.nextNodes {
        if node.value == values[place]:
            if place + 1 == len(values):
                return true
            return node.find(values, place + 1)
    }
}
