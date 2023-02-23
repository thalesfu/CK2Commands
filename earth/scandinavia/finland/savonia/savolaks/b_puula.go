package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉PuulaBarony struct {
	feud.BaseBarony
}

var BPuula普拉 feud.Barony = &普拉PuulaBarony{}

func init() {
	f := BPuula普拉.(*普拉PuulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puula",
		TitleName: "普拉",
		TitleCode: "b_puula",
	}
}
