package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5987] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5987][125987] = People_125987
	HistoryPeopleMap[5987][145987] = People_145987
}
