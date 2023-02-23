package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃德萨EdessaBarony struct {
	feud.BaseBarony
}

var BEdessa埃德萨 feud.Barony = &埃德萨EdessaBarony{}

func init() {
	f := BEdessa埃德萨.(*埃德萨EdessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edessa",
		TitleName: "埃德萨",
		TitleCode: "b_edessa",
	}
}
