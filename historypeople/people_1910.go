package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1910] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1910][71910] = People_71910
	HistoryPeopleMap[1910][161910] = People_161910
	HistoryPeopleMap[1910][191910] = People_191910
	HistoryPeopleMap[1910][461910] = People_461910
}
