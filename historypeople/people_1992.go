package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1992] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1992][71992] = People_71992
	HistoryPeopleMap[1992][191992] = People_191992
	HistoryPeopleMap[1992][451992] = People_451992
}
