package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6100] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6100][6100] = People_6100
	HistoryPeopleMap[6100][96100] = People_96100
	HistoryPeopleMap[6100][146100] = People_146100
	HistoryPeopleMap[6100][166100] = People_166100
	HistoryPeopleMap[6100][256100] = People_256100
	HistoryPeopleMap[6100][476100] = People_476100
}
