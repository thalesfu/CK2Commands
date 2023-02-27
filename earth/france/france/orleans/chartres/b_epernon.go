package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃佩尔农EpernonBarony struct {
	feud.BaseBarony
}

var BEpernon埃佩尔农 feud.Barony = &埃佩尔农EpernonBarony{}

func init() {
    f := BEpernon埃佩尔农.(*埃佩尔农EpernonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "epernon",
		TitleName: "埃佩尔农",
		TitleCode: "b_epernon",
	}
}
