package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒囊MangnangBarony struct {
	feud.BaseBarony
}

var BMangnang芒囊 feud.Barony = &芒囊MangnangBarony{}

func init() {
    f := BMangnang芒囊.(*芒囊MangnangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangnang",
		TitleName: "芒囊",
		TitleCode: "b_mangnang",
	}
}
