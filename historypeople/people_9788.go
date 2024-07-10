package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9788] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9788][159788] = People_159788
}
