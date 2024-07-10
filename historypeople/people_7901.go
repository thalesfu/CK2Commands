package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7901][7901] = People_7901
	HistoryPeopleMap[7901][167901] = People_167901
	HistoryPeopleMap[7901][247901] = People_247901
}
