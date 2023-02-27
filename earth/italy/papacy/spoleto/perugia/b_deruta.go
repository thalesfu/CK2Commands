package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德鲁塔DerutaBarony struct {
	feud.BaseBarony
}

var BDeruta德鲁塔 feud.Barony = &德鲁塔DerutaBarony{}

func init() {
    f := BDeruta德鲁塔.(*德鲁塔DerutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deruta",
		TitleName: "德鲁塔",
		TitleCode: "b_deruta",
	}
}
