package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[950] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[950][950] = People_950
	HistoryPeopleMap[950][30950] = People_30950
	HistoryPeopleMap[950][70950] = People_70950
	HistoryPeopleMap[950][260950] = People_260950
}
