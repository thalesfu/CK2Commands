package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉PulaBarony struct {
	feud.BaseBarony
}

var BPula普拉 feud.Barony = &普拉PulaBarony{}

func init() {
	f := BPula普拉.(*普拉PulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pula",
		TitleName: "普拉",
		TitleCode: "b_pula",
	}
}
