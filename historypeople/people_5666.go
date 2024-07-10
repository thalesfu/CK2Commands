package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5666][145666] = People_145666
}
