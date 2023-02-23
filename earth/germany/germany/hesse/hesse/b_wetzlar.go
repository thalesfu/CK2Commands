package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦茨拉尔WetzlarBarony struct {
	feud.BaseBarony
}

var BWetzlar韦茨拉尔 feud.Barony = &韦茨拉尔WetzlarBarony{}

func init() {
	f := BWetzlar韦茨拉尔.(*韦茨拉尔WetzlarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wetzlar",
		TitleName: "韦茨拉尔",
		TitleCode: "b_wetzlar",
	}
}
