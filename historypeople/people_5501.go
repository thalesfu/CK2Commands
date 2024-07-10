package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5501] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5501][125501] = People_125501
	HistoryPeopleMap[5501][215501] = People_215501
	HistoryPeopleMap[5501][455501] = People_455501
	HistoryPeopleMap[5501][465501] = People_465501
}
