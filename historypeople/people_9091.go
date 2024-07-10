package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9091] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9091][159091] = People_159091
	HistoryPeopleMap[9091][189091] = People_189091
}
