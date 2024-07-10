package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3298] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3298][33298] = People_33298
	HistoryPeopleMap[3298][73298] = People_73298
	HistoryPeopleMap[3298][93298] = People_93298
}
