package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努恩古NunguBarony struct {
	feud.BaseBarony
}

var BNungu努恩古 feud.Barony = &努恩古NunguBarony{}

func init() {
    f := BNungu努恩古.(*努恩古NunguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nungu",
		TitleName: "努恩古",
		TitleCode: "b_nungu",
	}
}
