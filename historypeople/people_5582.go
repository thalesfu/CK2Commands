package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5582] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5582][205582] = People_205582
}
