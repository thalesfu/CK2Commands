package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武科瓦尔VukovarBarony struct {
	feud.BaseBarony
}

var BVukovar武科瓦尔 feud.Barony = &武科瓦尔VukovarBarony{}

func init() {
    f := BVukovar武科瓦尔.(*武科瓦尔VukovarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vukovar",
		TitleName: "武科瓦尔",
		TitleCode: "b_vukovar",
	}
}
