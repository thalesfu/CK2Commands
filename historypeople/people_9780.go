package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9780] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9780][159780] = People_159780
}
