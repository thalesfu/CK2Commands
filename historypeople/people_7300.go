package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7300][7300] = People_7300
	HistoryPeopleMap[7300][157300] = People_157300
	HistoryPeopleMap[7300][167300] = People_167300
	HistoryPeopleMap[7300][217300] = People_217300
	HistoryPeopleMap[7300][247300] = People_247300
}
