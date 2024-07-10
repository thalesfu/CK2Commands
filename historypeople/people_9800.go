package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9800][159800] = People_159800
}
