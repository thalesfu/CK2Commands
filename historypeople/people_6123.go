package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6123] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6123][146123] = People_146123
	HistoryPeopleMap[6123][166123] = People_166123
	HistoryPeopleMap[6123][256123] = People_256123
}
