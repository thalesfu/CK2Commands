package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9810] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9810][9810] = People_9810
	HistoryPeopleMap[9810][159810] = People_159810
}
