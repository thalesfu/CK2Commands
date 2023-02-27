package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列科普PerekopBarony struct {
	feud.BaseBarony
}

var BPerekop佩列科普 feud.Barony = &佩列科普PerekopBarony{}

func init() {
    f := BPerekop佩列科普.(*佩列科普PerekopBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perekop",
		TitleName: "佩列科普",
		TitleCode: "b_perekop",
	}
}
