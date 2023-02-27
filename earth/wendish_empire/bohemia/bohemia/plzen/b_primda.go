package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普日姆达PrimdaBarony struct {
	feud.BaseBarony
}

var BPrimda普日姆达 feud.Barony = &普日姆达PrimdaBarony{}

func init() {
    f := BPrimda普日姆达.(*普日姆达PrimdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "primda",
		TitleName: "普日姆达",
		TitleCode: "b_primda",
	}
}
