package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丕林PerimBarony struct {
	feud.BaseBarony
}

var BPerim丕林 feud.Barony = &丕林PerimBarony{}

func init() {
	f := BPerim丕林.(*丕林PerimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perim",
		TitleName: "丕林",
		TitleCode: "b_perim",
	}
}
