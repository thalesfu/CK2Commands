package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺伊贝格NeubergBarony struct {
	feud.BaseBarony
}

var BNeuberg诺伊贝格 feud.Barony = &诺伊贝格NeubergBarony{}

func init() {
    f := BNeuberg诺伊贝格.(*诺伊贝格NeubergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neuberg",
		TitleName: "诺伊贝格",
		TitleCode: "b_neuberg",
	}
}
