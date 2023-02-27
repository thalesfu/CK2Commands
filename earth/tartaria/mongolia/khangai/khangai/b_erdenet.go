package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 额尔登特ErdenetBarony struct {
	feud.BaseBarony
}

var BErdenet额尔登特 feud.Barony = &额尔登特ErdenetBarony{}

func init() {
    f := BErdenet额尔登特.(*额尔登特ErdenetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erdenet",
		TitleName: "额尔登特",
		TitleCode: "b_erdenet",
	}
}
