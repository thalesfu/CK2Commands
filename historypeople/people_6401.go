package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6401] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6401][6401] = People_6401
	HistoryPeopleMap[6401][166401] = People_166401
}
