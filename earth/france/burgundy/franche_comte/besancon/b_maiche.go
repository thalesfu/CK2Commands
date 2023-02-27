package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈什MaicheBarony struct {
	feud.BaseBarony
}

var BMaiche迈什 feud.Barony = &迈什MaicheBarony{}

func init() {
    f := BMaiche迈什.(*迈什MaicheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maiche",
		TitleName: "迈什",
		TitleCode: "b_maiche",
	}
}
