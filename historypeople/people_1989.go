package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1989] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1989][71989] = People_71989
	HistoryPeopleMap[1989][191989] = People_191989
}
