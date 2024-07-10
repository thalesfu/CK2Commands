package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6300][166300] = People_166300
	HistoryPeopleMap[6300][476300] = People_476300
}
