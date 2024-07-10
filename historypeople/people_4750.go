package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4750] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4750][74750] = People_74750
	HistoryPeopleMap[4750][204750] = People_204750
}
