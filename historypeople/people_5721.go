package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5721] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5721][125721] = People_125721
	HistoryPeopleMap[5721][145721] = People_145721
}
