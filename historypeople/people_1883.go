package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1883] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1883][31883] = People_31883
	HistoryPeopleMap[1883][71883] = People_71883
	HistoryPeopleMap[1883][191883] = People_191883
}
