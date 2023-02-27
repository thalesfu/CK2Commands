package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什亚ShiyaBarony struct {
	feud.BaseBarony
}

var BShiya什亚 feud.Barony = &什亚ShiyaBarony{}

func init() {
    f := BShiya什亚.(*什亚ShiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiya",
		TitleName: "什亚",
		TitleCode: "b_shiya",
	}
}
