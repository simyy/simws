package dict

import (
    "log"
    "github.com/yxd123/simWS/tire"
)

type Dict struct {
    tire *tire
    maxLength int
}

func NewDict() *Dict {
    return &Dict{
        tire: tire.New(),
        maxLength: 0
    }
}

func (d *Dict) Add(word string, freq int) {
    if d.Find([]byte(word)) {
        return
    }

    res := d.Find([]byte(word))
    if res != nil {
        t.freq.AddFreq(freq)
        return
    }

    d.Add([]byte(word), freq)
}

func (d *Dict) Find(word string) int {
    res := d.tire.Find(d)
    if res != nil {
        return res.GetFreq()
    }

    return 0
}

func (d *Dict) Load(word string) {
    
}
