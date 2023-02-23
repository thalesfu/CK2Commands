package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎列西耶ZaleseBarony struct {
	feud.BaseBarony
}

var BZalese扎列西耶 feud.Barony = &扎列西耶ZaleseBarony{}

func init() {
	f := BZalese扎列西耶.(*扎列西耶ZaleseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zalese",
		TitleName: "扎列西耶",
		TitleCode: "b_zalese",
	}
}
