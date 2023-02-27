package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜尔KizilBarony struct {
	feud.BaseBarony
}

var BKizil克孜尔 feud.Barony = &克孜尔KizilBarony{}

func init() {
    f := BKizil克孜尔.(*克孜尔KizilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizil",
		TitleName: "克孜尔",
		TitleCode: "b_kizil",
	}
}
