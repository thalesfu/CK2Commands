package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5401] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5401][125401] = People_125401
}
