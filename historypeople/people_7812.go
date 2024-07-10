package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7812] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7812][7812] = People_7812
	HistoryPeopleMap[7812][167812] = People_167812
	HistoryPeopleMap[7812][247812] = People_247812
}
