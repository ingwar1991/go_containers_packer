package packer

import (
	"maps"
	"math"
)


type (
    containerSize int
    containersCnt int

    goodsSum int
)

type approach  struct {
    id int 
    containers map[containerSize]int
    storedGoods goodsSum 
    containersTotal containersCnt 
}

func (a *approach) ID() int {
    return a.id
}

func (a *approach) Containers() map[containerSize]int {
    return a.containers
}

func (a *approach) StoredGoods() goodsSum {
    return a.storedGoods
}

func (a *approach) ContainersTotal() containersCnt {
    return a.containersTotal
}

func (a *approach) AddContainer(size containerSize) {
    if _, ok := a.containers[size]; !ok {
        a.containers[size] = 0
    }

    a.containers[size]++
    a.storedGoods += goodsSum(size)
    a.containersTotal++
}

func (a *approach) IsCompleted(sum goodsSum) bool {
    return a.StoredGoods() >= sum
}

func (a *approach) IsBetter(sum goodsSum, apprch *approach) bool {
    aAbs := math.Abs(float64(sum - a.StoredGoods()))
    apprchAbs := math.Abs(float64(sum - apprch.StoredGoods()))

    if aAbs < apprchAbs {
        return true
    }

    if aAbs == apprchAbs || (a.IsCompleted(sum) && !apprch.IsCompleted(sum)) {
        if a.ContainersTotal() < apprch.ContainersTotal() {
            return true
        }
    }

    return false
}

func (a *approach) NewApproachFromExisting(apprchID *int, size containerSize) *approach {
    newContainers := make(map[containerSize]int, len(a.containers))
    maps.Copy(newContainers, a.containers)

    *apprchID++
    newapproach := approach{
        id: *apprchID, 
        containers: newContainers, 
        storedGoods: a.storedGoods, 
        containersTotal: a.containersTotal, 
    }
    newapproach.AddContainer(size)

    return &newapproach
}

func NewApproach(apprchID *int, size containerSize) *approach {
    *apprchID++
    return &approach{
        id: *apprchID, 
        containers: map[containerSize]int{
            size: 1,
        }, 
        storedGoods: goodsSum(size), 
        containersTotal: 1, 
    }
}
