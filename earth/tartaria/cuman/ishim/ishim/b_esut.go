package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃苏特EsutBarony struct {
	feud.BaseBarony
}

var BEsut埃苏特 feud.Barony = &埃苏特EsutBarony{}

func init() {
    f := BEsut埃苏特.(*埃苏特EsutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esut",
		TitleName: "埃苏特",
		TitleCode: "b_esut",
	}
}
