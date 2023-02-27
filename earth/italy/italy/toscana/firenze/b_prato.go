package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉托PratoBarony struct {
	feud.BaseBarony
}

var BPrato普拉托 feud.Barony = &普拉托PratoBarony{}

func init() {
    f := BPrato普拉托.(*普拉托PratoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prato",
		TitleName: "普拉托",
		TitleCode: "b_prato",
	}
}
