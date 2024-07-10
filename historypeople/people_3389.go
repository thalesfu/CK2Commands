package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3389] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3389][33389] = People_33389
	HistoryPeopleMap[3389][73389] = People_73389
	HistoryPeopleMap[3389][83389] = People_83389
	HistoryPeopleMap[3389][93389] = People_93389
}
