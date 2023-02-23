package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利夫内LivnyBarony struct {
	feud.BaseBarony
}

var BLivny利夫内 feud.Barony = &利夫内LivnyBarony{}

func init() {
	f := BLivny利夫内.(*利夫内LivnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "livny",
		TitleName: "利夫内",
		TitleCode: "b_livny",
	}
}
