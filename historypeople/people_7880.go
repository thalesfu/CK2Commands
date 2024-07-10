package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7880] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7880][7880] = People_7880
	HistoryPeopleMap[7880][167880] = People_167880
	HistoryPeopleMap[7880][247880] = People_247880
}
