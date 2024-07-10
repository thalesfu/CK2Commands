package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1978] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1978][71978] = People_71978
	HistoryPeopleMap[1978][161978] = People_161978
	HistoryPeopleMap[1978][191978] = People_191978
}
