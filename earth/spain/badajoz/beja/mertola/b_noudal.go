package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺达尔NoudalBarony struct {
	feud.BaseBarony
}

var BNoudal诺达尔 feud.Barony = &诺达尔NoudalBarony{}

func init() {
    f := BNoudal诺达尔.(*诺达尔NoudalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noudal",
		TitleName: "诺达尔",
		TitleCode: "b_noudal",
	}
}
