package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5265] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5265][145265] = People_145265
	HistoryPeopleMap[5265][205265] = People_205265
}
