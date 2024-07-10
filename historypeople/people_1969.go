package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1969] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1969][71969] = People_71969
	HistoryPeopleMap[1969][161969] = People_161969
	HistoryPeopleMap[1969][191969] = People_191969
}
