package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6801][6801] = People_6801
	HistoryPeopleMap[6801][166801] = People_166801
	HistoryPeopleMap[6801][206801] = People_206801
}
