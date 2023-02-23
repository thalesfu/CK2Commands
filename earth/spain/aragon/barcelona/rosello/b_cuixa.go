package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库沙CuixaBarony struct {
	feud.BaseBarony
}

var BCuixa库沙 feud.Barony = &库沙CuixaBarony{}

func init() {
	f := BCuixa库沙.(*库沙CuixaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuixa",
		TitleName: "库沙",
		TitleCode: "b_cuixa",
	}
}
