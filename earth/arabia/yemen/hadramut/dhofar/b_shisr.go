package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什斯尔ShisrBarony struct {
	feud.BaseBarony
}

var BShisr什斯尔 feud.Barony = &什斯尔ShisrBarony{}

func init() {
    f := BShisr什斯尔.(*什斯尔ShisrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shisr",
		TitleName: "什斯尔",
		TitleCode: "b_shisr",
	}
}
