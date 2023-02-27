package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔莱MorlaixBarony struct {
	feud.BaseBarony
}

var BMorlaix莫尔莱 feud.Barony = &莫尔莱MorlaixBarony{}

func init() {
    f := BMorlaix莫尔莱.(*莫尔莱MorlaixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morlaix",
		TitleName: "莫尔莱",
		TitleCode: "b_morlaix",
	}
}
