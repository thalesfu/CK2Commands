package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内穆尔NemoursBarony struct {
	feud.BaseBarony
}

var BNemours内穆尔 feud.Barony = &内穆尔NemoursBarony{}

func init() {
	f := BNemours内穆尔.(*内穆尔NemoursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemours",
		TitleName: "内穆尔",
		TitleCode: "b_nemours",
	}
}
