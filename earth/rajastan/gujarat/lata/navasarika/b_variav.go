package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐里瓦VariavBarony struct {
	feud.BaseBarony
}

var BVariav伐里瓦 feud.Barony = &伐里瓦VariavBarony{}

func init() {
	f := BVariav伐里瓦.(*伐里瓦VariavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "variav",
		TitleName: "伐里瓦",
		TitleCode: "b_variav",
	}
}
