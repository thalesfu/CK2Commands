package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽梅诺斯ZemenosBarony struct {
	feud.BaseBarony
}

var BZemenos泽梅诺斯 feud.Barony = &泽梅诺斯ZemenosBarony{}

func init() {
    f := BZemenos泽梅诺斯.(*泽梅诺斯ZemenosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zemenos",
		TitleName: "泽梅诺斯",
		TitleCode: "b_zemenos",
	}
}
