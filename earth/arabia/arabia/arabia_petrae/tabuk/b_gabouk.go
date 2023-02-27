package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加布克GaboukBarony struct {
	feud.BaseBarony
}

var BGabouk加布克 feud.Barony = &加布克GaboukBarony{}

func init() {
    f := BGabouk加布克.(*加布克GaboukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabouk",
		TitleName: "加布克",
		TitleCode: "b_gabouk",
	}
}
