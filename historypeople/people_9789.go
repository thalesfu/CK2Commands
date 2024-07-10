package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9789][159789] = People_159789
}
