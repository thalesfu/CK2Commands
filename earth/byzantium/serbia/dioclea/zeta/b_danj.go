package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尼DanjBarony struct {
	feud.BaseBarony
}

var BDanj达尼 feud.Barony = &达尼DanjBarony{}

func init() {
	f := BDanj达尼.(*达尼DanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danj",
		TitleName: "达尼",
		TitleCode: "b_danj",
	}
}
