package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6868] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6868][6868] = People_6868
	HistoryPeopleMap[6868][166868] = People_166868
}
