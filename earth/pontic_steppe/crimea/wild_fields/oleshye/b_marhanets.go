package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔加涅茨MarhanetsBarony struct {
	feud.BaseBarony
}

var BMarhanets马尔加涅茨 feud.Barony = &马尔加涅茨MarhanetsBarony{}

func init() {
    f := BMarhanets马尔加涅茨.(*马尔加涅茨MarhanetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marhanets",
		TitleName: "马尔加涅茨",
		TitleCode: "b_marhanets",
	}
}
