package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1871] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1871][31871] = People_31871
	HistoryPeopleMap[1871][71871] = People_71871
	HistoryPeopleMap[1871][81871] = People_81871
	HistoryPeopleMap[1871][191871] = People_191871
}
