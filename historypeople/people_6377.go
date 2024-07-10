package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6377] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6377][166377] = People_166377
}
