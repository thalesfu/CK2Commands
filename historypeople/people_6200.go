package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6200][6200] = People_6200
	HistoryPeopleMap[6200][146200] = People_146200
	HistoryPeopleMap[6200][166200] = People_166200
	HistoryPeopleMap[6200][476200] = People_476200
}
