package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃莱亚ElaiaBarony struct {
	feud.BaseBarony
}

var BElaia埃莱亚 feud.Barony = &埃莱亚ElaiaBarony{}

func init() {
    f := BElaia埃莱亚.(*埃莱亚ElaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elaia",
		TitleName: "埃莱亚",
		TitleCode: "b_elaia",
	}
}
