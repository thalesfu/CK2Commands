package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7500] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7500][7500] = People_7500
	HistoryPeopleMap[7500][107500] = People_107500
	HistoryPeopleMap[7500][167500] = People_167500
	HistoryPeopleMap[7500][207500] = People_207500
	HistoryPeopleMap[7500][217500] = People_217500
	HistoryPeopleMap[7500][247500] = People_247500
	HistoryPeopleMap[7500][457500] = People_457500
	HistoryPeopleMap[7500][467500] = People_467500
}
