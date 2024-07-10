package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4666][74666] = People_74666
	HistoryPeopleMap[4666][204666] = People_204666
}
