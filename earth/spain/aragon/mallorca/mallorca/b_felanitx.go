package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费拉尼奇FelanitxBarony struct {
	feud.BaseBarony
}

var BFelanitx费拉尼奇 feud.Barony = &费拉尼奇FelanitxBarony{}

func init() {
    f := BFelanitx费拉尼奇.(*费拉尼奇FelanitxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "felanitx",
		TitleName: "费拉尼奇",
		TitleCode: "b_felanitx",
	}
}
