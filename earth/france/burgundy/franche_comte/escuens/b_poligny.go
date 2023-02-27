package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利尼PolignyBarony struct {
	feud.BaseBarony
}

var BPoligny波利尼 feud.Barony = &波利尼PolignyBarony{}

func init() {
    f := BPoligny波利尼.(*波利尼PolignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poligny",
		TitleName: "波利尼",
		TitleCode: "b_poligny",
	}
}
