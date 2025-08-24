package packer

import (
    "fmt"
	"sort"
)


func Pack(containersList []int, goodsCnt int) (map[int]int, error) {
    if len(containersList) < 1 {
        return nil, fmt.Errorf("No containersList transmitted: %v", containersList)
    }
    for _, container := range containersList {
        if container < 1 {
            return nil, fmt.Errorf("Container size has to be positive int: %v", containersList)
        }
    }
    if goodsCnt < 1 {
        return nil, fmt.Errorf("Goods size has to be positive int: %v", goodsCnt)
    }

    sort.Ints(containersList)
    containers := make([]containerSize, len(containersList))
    for i, size := range containersList {
        containers[i] = containerSize(size)
    }

    res := make(map[int]int)
    queue := Queue{
        lastApproachID: new(int),
        approachSearch: make(map[goodsSum]containersCnt),    
        approaches: make(map[int]*approach),
        candidates: make(map[int]*approach),
    } 

    var bestApproach *approach
    for _, size := range containers {
        approach := NewApproach(queue.LastApproachID(), containerSize(size))
        bestApproach = approach

        queue.AddCandidate(approach) 
    }
    queue.PromoteCandidates()

    gSum := goodsSum(goodsCnt)
    for {
        for _, approach := range queue.Approaches() {
            if approach.IsCompleted(gSum) {
                if approach.IsBetter(gSum, bestApproach) {
                    bestApproach = approach
                }

                continue
            }

            for _, size := range containers {
                newApproach := approach.NewApproachFromExisting(queue.LastApproachID(), containerSize(size))
                
                if cnt, ok := queue.Search(newApproach.StoredGoods()); ok && cnt <= newApproach.ContainersTotal() {
                    newApproach = nil
                    continue
                }

                queue.AddCandidate(newApproach)

                if newApproach.IsBetter(gSum, bestApproach) {
                    bestApproach = newApproach
                }
            }
        }

        queue.PromoteCandidates()

        if queue.IsFinished(gSum) {
            break
        }
    }

    for size, cnt := range bestApproach.Containers() {
        res[int(size)] = cnt
    }

    return res, nil
}
