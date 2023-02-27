package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里布尼克RybnikBarony struct {
	feud.BaseBarony
}

var BRybnik里布尼克 feud.Barony = &里布尼克RybnikBarony{}

func init() {
    f := BRybnik里布尼克.(*里布尼克RybnikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rybnik",
		TitleName: "里布尼克",
		TitleCode: "b_rybnik",
	}
}
