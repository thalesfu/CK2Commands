package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃提梨ChanderiBarony struct {
	feud.BaseBarony
}

var BChanderi旃提梨 feud.Barony = &旃提梨ChanderiBarony{}

func init() {
    f := BChanderi旃提梨.(*旃提梨ChanderiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chanderi",
		TitleName: "旃提梨",
		TitleCode: "b_chanderi",
	}
}
