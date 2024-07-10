package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[912] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[912][912] = People_912
	HistoryPeopleMap[912][30912] = People_30912
	HistoryPeopleMap[912][70912] = People_70912
	HistoryPeopleMap[912][260912] = People_260912
}
