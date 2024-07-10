package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1898] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1898][191898] = People_191898
}
