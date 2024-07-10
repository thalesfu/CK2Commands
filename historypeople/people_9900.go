package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9900][159900] = People_159900
}
