package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那揭罗波罗迦罗NagarparkarBarony struct {
	feud.BaseBarony
}

var BNagarparkar那揭罗波罗迦罗 feud.Barony = &那揭罗波罗迦罗NagarparkarBarony{}

func init() {
    f := BNagarparkar那揭罗波罗迦罗.(*那揭罗波罗迦罗NagarparkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagarparkar",
		TitleName: "那揭罗波罗迦罗",
		TitleCode: "b_nagarparkar",
	}
}
