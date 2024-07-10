package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1920] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1920][71920] = People_71920
	HistoryPeopleMap[1920][161920] = People_161920
	HistoryPeopleMap[1920][191920] = People_191920
	HistoryPeopleMap[1920][461920] = People_461920
}
