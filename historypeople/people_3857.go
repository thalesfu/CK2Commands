package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3857] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3857][73857] = People_73857
}
