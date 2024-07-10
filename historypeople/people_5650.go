package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5650] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5650][5650] = People_5650
	HistoryPeopleMap[5650][145650] = People_145650
	HistoryPeopleMap[5650][205650] = People_205650
	HistoryPeopleMap[5650][455650] = People_455650
}
