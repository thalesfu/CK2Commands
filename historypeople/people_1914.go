package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1914] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1914][161914] = People_161914
	HistoryPeopleMap[1914][191914] = People_191914
	HistoryPeopleMap[1914][461914] = People_461914
}
