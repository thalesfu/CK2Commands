package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5772] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5772][145772] = People_145772
}
