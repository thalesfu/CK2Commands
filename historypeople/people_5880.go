package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5880] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5880][145880] = People_145880
	HistoryPeopleMap[5880][205880] = People_205880
}
