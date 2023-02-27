package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德任LadyzhynBarony struct {
	feud.BaseBarony
}

var BLadyzhyn拉德任 feud.Barony = &拉德任LadyzhynBarony{}

func init() {
    f := BLadyzhyn拉德任.(*拉德任LadyzhynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ladyzhyn",
		TitleName: "拉德任",
		TitleCode: "b_ladyzhyn",
	}
}
