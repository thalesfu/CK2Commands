package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[750] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[750][30750] = People_30750
	HistoryPeopleMap[750][70750] = People_70750
	HistoryPeopleMap[750][180750] = People_180750
	HistoryPeopleMap[750][260750] = People_260750
	HistoryPeopleMap[750][450750] = People_450750
}
