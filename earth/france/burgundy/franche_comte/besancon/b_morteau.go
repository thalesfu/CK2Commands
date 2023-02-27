package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔托MorteauBarony struct {
	feud.BaseBarony
}

var BMorteau莫尔托 feud.Barony = &莫尔托MorteauBarony{}

func init() {
    f := BMorteau莫尔托.(*莫尔托MorteauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morteau",
		TitleName: "莫尔托",
		TitleCode: "b_morteau",
	}
}
