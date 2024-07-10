package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1975] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1975][71975] = People_71975
	HistoryPeopleMap[1975][161975] = People_161975
	HistoryPeopleMap[1975][191975] = People_191975
	HistoryPeopleMap[1975][451975] = People_451975
}
