package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[799] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[799][20799] = People_20799
	HistoryPeopleMap[799][30799] = People_30799
	HistoryPeopleMap[799][70799] = People_70799
	HistoryPeopleMap[799][250799] = People_250799
}
