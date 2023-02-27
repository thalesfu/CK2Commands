package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗斯齐亚KnudskerBarony struct {
	feud.BaseBarony
}

var BKnudsker克罗斯齐亚 feud.Barony = &克罗斯齐亚KnudskerBarony{}

func init() {
    f := BKnudsker克罗斯齐亚.(*克罗斯齐亚KnudskerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knudsker",
		TitleName: "克罗斯齐亚",
		TitleCode: "b_knudsker",
	}
}
