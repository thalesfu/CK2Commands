package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6502] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6502][166502] = People_166502
	HistoryPeopleMap[6502][186502] = People_186502
	HistoryPeopleMap[6502][206502] = People_206502
	HistoryPeopleMap[6502][466502] = People_466502
	HistoryPeopleMap[6502][476502] = People_476502
}
