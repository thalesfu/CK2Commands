package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3280] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3280][73280] = People_73280
	HistoryPeopleMap[3280][83280] = People_83280
	HistoryPeopleMap[3280][93280] = People_93280
}
