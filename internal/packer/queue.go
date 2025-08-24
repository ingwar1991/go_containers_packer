package packer

import (
	"maps"
)


type Queue struct {
    lastApproachID *int
    approachSearch map[goodsSum]containersCnt
    approaches map[int]*approach
    candidates map[int]*approach
}

func (q *Queue) LastApproachID() *int {
    return q.lastApproachID
}

func (q *Queue) AddCandidate(approach *approach) {
    q.candidates[approach.ID()] = approach
    q.approachSearch[approach.StoredGoods()] = approach.ContainersTotal()
}

func (q *Queue) PromoteCandidates() {
    q.approaches = make(map[int]*approach, len(q.candidates))
    maps.Copy(q.approaches, q.candidates)

    q.candidates = make(map[int]*approach)
}

func (q *Queue) Remove(approach *approach) {
    delete(q.approaches, approach.ID())
    approach = nil
}

func (q *Queue) Approaches() map[int]*approach {
    return q.approaches
}

func (q *Queue) Search(sum goodsSum) (containersCnt, bool) {
    val, ok := q.approachSearch[sum] 
    return val, ok
}

func (q *Queue) IsFinished(goodSum goodsSum) bool {
    approachesCnt := len(q.approaches)
    for _, approach := range q.approaches {
        if approach.IsCompleted(goodSum) {
            approachesCnt--;
        }
    }

    return approachesCnt < 1 
}
