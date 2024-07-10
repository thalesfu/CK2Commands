package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1897] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1897][191897] = People_191897
}
