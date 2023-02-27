package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥本罗AabenraaBarony struct {
	feud.BaseBarony
}

var BAabenraa奥本罗 feud.Barony = &奥本罗AabenraaBarony{}

func init() {
    f := BAabenraa奥本罗.(*奥本罗AabenraaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aabenraa",
		TitleName: "奥本罗",
		TitleCode: "b_aabenraa",
	}
}
