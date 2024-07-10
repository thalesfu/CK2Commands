package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2950] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2950][32950] = People_32950
	HistoryPeopleMap[2950][72950] = People_72950
	HistoryPeopleMap[2950][142950] = People_142950
	HistoryPeopleMap[2950][212950] = People_212950
}
