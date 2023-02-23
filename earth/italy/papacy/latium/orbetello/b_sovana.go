package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索瓦纳SovanaBarony struct {
	feud.BaseBarony
}

var BSovana索瓦纳 feud.Barony = &索瓦纳SovanaBarony{}

func init() {
	f := BSovana索瓦纳.(*索瓦纳SovanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sovana",
		TitleName: "索瓦纳",
		TitleCode: "b_sovana",
	}
}
