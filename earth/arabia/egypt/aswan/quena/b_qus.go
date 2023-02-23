package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古斯QusBarony struct {
	feud.BaseBarony
}

var BQus古斯 feud.Barony = &古斯QusBarony{}

func init() {
	f := BQus古斯.(*古斯QusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qus",
		TitleName: "古斯",
		TitleCode: "b_qus",
	}
}
