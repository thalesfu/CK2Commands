package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗奢罗波RajrappaBarony struct {
	feud.BaseBarony
}

var BRajrappa罗奢罗波 feud.Barony = &罗奢罗波RajrappaBarony{}

func init() {
    f := BRajrappa罗奢罗波.(*罗奢罗波RajrappaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajrappa",
		TitleName: "罗奢罗波",
		TitleCode: "b_rajrappa",
	}
}
