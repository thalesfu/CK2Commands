package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿悉言AxiyanBarony struct {
	feud.BaseBarony
}

var BAxiyan阿悉言 feud.Barony = &阿悉言AxiyanBarony{}

func init() {
    f := BAxiyan阿悉言.(*阿悉言AxiyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "axiyan",
		TitleName: "阿悉言",
		TitleCode: "b_axiyan",
	}
}
