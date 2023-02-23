package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里瓦RivaBarony struct {
	feud.BaseBarony
}

var BRiva里瓦 feud.Barony = &里瓦RivaBarony{}

func init() {
	f := BRiva里瓦.(*里瓦RivaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "riva",
		TitleName: "里瓦",
		TitleCode: "b_riva",
	}
}
