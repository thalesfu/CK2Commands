package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7778] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7778][7778] = People_7778
	HistoryPeopleMap[7778][167778] = People_167778
	HistoryPeopleMap[7778][247778] = People_247778
}
