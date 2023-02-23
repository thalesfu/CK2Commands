package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯拉蒂纳SlatinaBarony struct {
	feud.BaseBarony
}

var BSlatina斯拉蒂纳 feud.Barony = &斯拉蒂纳SlatinaBarony{}

func init() {
	f := BSlatina斯拉蒂纳.(*斯拉蒂纳SlatinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slatina",
		TitleName: "斯拉蒂纳",
		TitleCode: "b_slatina",
	}
}
