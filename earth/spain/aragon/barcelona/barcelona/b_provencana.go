package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗文萨纳ProvencanaBarony struct {
	feud.BaseBarony
}

var BProvencana普罗文萨纳 feud.Barony = &普罗文萨纳ProvencanaBarony{}

func init() {
	f := BProvencana普罗文萨纳.(*普罗文萨纳ProvencanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "provencana",
		TitleName: "普罗文萨纳",
		TitleCode: "b_provencana",
	}
}
