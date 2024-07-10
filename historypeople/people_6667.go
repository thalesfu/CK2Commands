package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6667] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6667][166667] = People_166667
	HistoryPeopleMap[6667][206667] = People_206667
}
