package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔加兹DaregazBarony struct {
	feud.BaseBarony
}

var BDaregaz达尔加兹 feud.Barony = &达尔加兹DaregazBarony{}

func init() {
    f := BDaregaz达尔加兹.(*达尔加兹DaregazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daregaz",
		TitleName: "达尔加兹",
		TitleCode: "b_daregaz",
	}
}
