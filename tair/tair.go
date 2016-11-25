package tair

import (
    "log"
)

type Tair struct {
    root Node
    count int
}

func New() Tair {
    return Tair{root: NewNode('0'), count: 0}
}

func (t *Tair) Count() int {
    return t.count
}

func (t *Tair) Find(values []byte) bool {
    if t.count == 0 {
        return false
    }

    return t.root.find(values, 0)
}

func (t *Tair) Add(values []byte) {
    if t.Find(values) {
        return
    }
    t.root.add(values, 0)
    t.count += 1
}

func (t *Tair) Del(values []byte) {
    if t.Find(values) == true {
        t.root.del(values, 0)
        t.count -= 1
    }
}

type Node struct {
    value byte
    nextNodes []Node
    nextNums int
}

func NewNode(value byte) Node {
    return Node{
        value: value,
        nextNodes: nil,
        nextNums: 0}
}

func (t *Node) add(values []byte, place int) {
    if place >= len(values) {
        return
    }

    if t.nextNodes == nil {
        log.Println("new nextNodes")
        t.nextNodes = make([]Node, 0, 26)
        t.nextNums = 0
    }

    if t.nextNums > 0 {
        for _, node := range t.nextNodes {
            if node.value == values[place] {
                log.Println("duplicate byte:", string(node.value))
                node.add(values, place + 1)
                return
            }
        }
    }
    node := NewNode(values[place])
    node.add(values, place + 1)
    t.nextNodes = append(t.nextNodes, node)
    t.nextNums += 1
    log.Println("add new Node:", string(node.value), t.nextNums)
}

func (t *Node) del(values []byte, place int) {
    if t.nextNums == 0 || place >= len(values) {
        return
    }

    for i, node := range t.nextNodes {
        if node.value == values[place] {
            node.del(values, place + 1)
            if node.nextNums <= 0 && place == len(values) - 1 {
                t.nextNodes = append(t.nextNodes[:i], t.nextNodes[i+1:]...) // replace slice, use *list.List 
            }
            break
        }
        break
    }
}

func (t *Node) find(values []byte, place int) bool {
    if t == nil || t.nextNums == 0 {
        log.Println("can't cmp", string(values[place]), "nextNums:", string(t.value), t.nextNums)
        return false
    }

    for _, node := range t.nextNodes {
        log.Println("cmp:", string(node.value), string(values[place]))
        if node.value == values[place] {
            if place == len(values) - 1 {
                return true
            }
            return node.find(values, place + 1)
        }
    }

    return false
}
