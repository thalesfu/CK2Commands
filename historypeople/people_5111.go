package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5111] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5111][125111] = People_125111
	HistoryPeopleMap[5111][145111] = People_145111
	HistoryPeopleMap[5111][155111] = People_155111
	HistoryPeopleMap[5111][205111] = People_205111
	HistoryPeopleMap[5111][235111] = People_235111
}
