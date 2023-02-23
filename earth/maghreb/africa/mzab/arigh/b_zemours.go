package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽穆尔ZemoursBarony struct {
	feud.BaseBarony
}

var BZemours泽穆尔 feud.Barony = &泽穆尔ZemoursBarony{}

func init() {
	f := BZemours泽穆尔.(*泽穆尔ZemoursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zemours",
		TitleName: "泽穆尔",
		TitleCode: "b_zemours",
	}
}
