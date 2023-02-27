package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波卢提ParodiBarony struct {
	feud.BaseBarony
}

var BParodi波卢提 feud.Barony = &波卢提ParodiBarony{}

func init() {
    f := BParodi波卢提.(*波卢提ParodiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parodi",
		TitleName: "波卢提",
		TitleCode: "b_parodi",
	}
}
