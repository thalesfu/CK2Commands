package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5701] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5701][125701] = People_125701
	HistoryPeopleMap[5701][145701] = People_145701
	HistoryPeopleMap[5701][205701] = People_205701
}
