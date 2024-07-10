package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5890] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5890][145890] = People_145890
	HistoryPeopleMap[5890][205890] = People_205890
}
