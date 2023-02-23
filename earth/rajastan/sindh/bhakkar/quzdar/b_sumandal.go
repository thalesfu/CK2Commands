package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏曼多SumandalBarony struct {
	feud.BaseBarony
}

var BSumandal苏曼多 feud.Barony = &苏曼多SumandalBarony{}

func init() {
	f := BSumandal苏曼多.(*苏曼多SumandalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumandal",
		TitleName: "苏曼多",
		TitleCode: "b_sumandal",
	}
}
