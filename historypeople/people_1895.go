package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1895] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1895][191895] = People_191895
}
