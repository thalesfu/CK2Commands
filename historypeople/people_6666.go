package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6666][166666] = People_166666
	HistoryPeopleMap[6666][206666] = People_206666
}
