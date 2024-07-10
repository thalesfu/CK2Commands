package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4730] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4730][74730] = People_74730
	HistoryPeopleMap[4730][204730] = People_204730
}
