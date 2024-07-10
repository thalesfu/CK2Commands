package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6820] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6820][6820] = People_6820
	HistoryPeopleMap[6820][166820] = People_166820
}
