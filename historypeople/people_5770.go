package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5770] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5770][145770] = People_145770
	HistoryPeopleMap[5770][205770] = People_205770
	HistoryPeopleMap[5770][455770] = People_455770
}
