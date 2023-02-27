package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达默DammeBarony struct {
	feud.BaseBarony
}

var BDamme达默 feud.Barony = &达默DammeBarony{}

func init() {
    f := BDamme达默.(*达默DammeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damme",
		TitleName: "达默",
		TitleCode: "b_damme",
	}
}
