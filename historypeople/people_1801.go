package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1801][31801] = People_31801
	HistoryPeopleMap[1801][71801] = People_71801
	HistoryPeopleMap[1801][191801] = People_191801
	HistoryPeopleMap[1801][451801] = People_451801
	HistoryPeopleMap[1801][461801] = People_461801
}
