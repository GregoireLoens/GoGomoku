package ai

import (
	"container/list"
)

func Any(l *list.List, v Position) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(Position).X == v.X && e.Value.(Position).Y == v.Y {
			return true
		}
	}
	return false
}
