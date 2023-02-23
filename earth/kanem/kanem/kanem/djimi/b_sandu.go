package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑杜SanduBarony struct {
	feud.BaseBarony
}

var BSandu桑杜 feud.Barony = &桑杜SanduBarony{}

func init() {
	f := BSandu桑杜.(*桑杜SanduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandu",
		TitleName: "桑杜",
		TitleCode: "b_sandu",
	}
}
