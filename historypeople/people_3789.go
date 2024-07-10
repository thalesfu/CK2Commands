package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3789][33789] = People_33789
	HistoryPeopleMap[3789][73789] = People_73789
	HistoryPeopleMap[3789][83789] = People_83789
	HistoryPeopleMap[3789][93789] = People_93789
}
