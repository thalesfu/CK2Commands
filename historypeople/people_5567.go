package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5567] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5567][205567] = People_205567
}
