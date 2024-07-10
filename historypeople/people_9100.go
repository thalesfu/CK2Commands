package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9100] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9100][159100] = People_159100
	HistoryPeopleMap[9100][189100] = People_189100
	HistoryPeopleMap[9100][479100] = People_479100
}
