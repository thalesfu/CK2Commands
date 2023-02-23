package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内尔瓦NervaBarony struct {
	feud.BaseBarony
}

var BNerva内尔瓦 feud.Barony = &内尔瓦NervaBarony{}

func init() {
	f := BNerva内尔瓦.(*内尔瓦NervaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nerva",
		TitleName: "内尔瓦",
		TitleCode: "b_nerva",
	}
}
