package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加来CalaisBarony struct {
	feud.BaseBarony
}

var BCalais加来 feud.Barony = &加来CalaisBarony{}

func init() {
	f := BCalais加来.(*加来CalaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calais",
		TitleName: "加来",
		TitleCode: "b_calais",
	}
}
