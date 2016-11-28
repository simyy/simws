package tire

import (
    "log"
)

type Tire struct {
    root *Node
    count int
}

func New() Tire {
    return Tire{root: NewNode('0'), count: 0}
}

func (t *Tire) Count() int {
    return t.count
}

func (t *Tire) Find(values []byte) *Node {
    if t.count == 0 {
        return nil
    }

    return t.root.Find(values, 0)
}

func (t *Tire) Add(values []byte, freq int) {
    if t.Find(values) != nil {
        return
    }
    t.root.Add(values, 0, freq)
    t.count += 1
}

func (t *Tire) Del(values []byte) {
    if t.Find(values) != nil {
        t.root.Del(values, 0)
        t.count -= 1
    }
}

type Node struct {
    value     byte
    nextNodes []*Node
    nextNums  int
    freq      int
}

func NewNode(value byte) *Node {
    return &Node{
        value: value,
        nextNodes: nil,
        nextNums: 0,
        freq: 0}
}

func (t *Node) Add(values []byte, place int, freq int) {
    if place >= len(values) {
        t.freq = freq
        return
    }

    if t.nextNodes == nil {
        log.Println("new nextNodes")
        t.nextNodes = make([]*Node, 0, 26)
        t.nextNums = 0
    }

    if t.nextNums > 0 {
        for _, node := range t.nextNodes {
            if node.value == values[place] {
                log.Println("duplicate byte:", string(node.value))
                node.Add(values, place + 1, freq)
                return
            }
        }
    }
    node := NewNode(values[place])
    node.Add(values, place + 1, freq)
    t.nextNodes = append(t.nextNodes, node)
    t.nextNums += 1
    log.Println("add new Node:", string(node.value), t.nextNums)
}

func (t *Node) Del(values []byte, place int) {
    if t.nextNums == 0 || place >= len(values) {
        return
    }

    for i, node := range t.nextNodes {
        if node.value == values[place] {
            node.Del(values, place + 1)
            if node.nextNums <= 0 && place == len(values) - 1 {
                t.nextNodes = append(t.nextNodes[:i], t.nextNodes[i+1:]...) // replace slice, use *list.List 
            }
            break
        }
        break
    }
}

func (t *Node) Find(values []byte, place int) *Node {
    if t == nil || t.nextNums == 0 {
        log.Println("can't cmp", string(values[place]), "nextNums:", string(t.value), t.nextNums)
        return nil
    }

    for _, node := range t.nextNodes {
        log.Println("cmp:", string(node.value), string(values[place]))
        if node.value == values[place] {
            if place == len(values) - 1 {
                if node.freq > 0 {
                    return node
                }
                return nil
            }
            return node.Find(values, place + 1)
        }
    }

    return nil
}

func (t *Node) AddFreq(freq int) {
    t.freq += freq
}

func (t *Node) GetFreq() int {
    return t.freq
}
