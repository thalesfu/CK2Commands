package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9090] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9090][129090] = People_129090
	HistoryPeopleMap[9090][159090] = People_159090
	HistoryPeopleMap[9090][189090] = People_189090
	HistoryPeopleMap[9090][479090] = People_479090
}
