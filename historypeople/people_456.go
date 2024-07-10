package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[456] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[456][30456] = People_30456
	HistoryPeopleMap[456][190456] = People_190456
	HistoryPeopleMap[456][260456] = People_260456
}
