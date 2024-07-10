package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7060] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7060][7060] = People_7060
	HistoryPeopleMap[7060][247060] = People_247060
	HistoryPeopleMap[7060][487060] = People_487060
}
