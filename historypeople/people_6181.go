package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6181] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6181][6181] = People_6181
	HistoryPeopleMap[6181][146181] = People_146181
	HistoryPeopleMap[6181][166181] = People_166181
}
