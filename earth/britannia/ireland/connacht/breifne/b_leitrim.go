package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利特里姆LeitrimBarony struct {
	feud.BaseBarony
}

var BLeitrim利特里姆 feud.Barony = &利特里姆LeitrimBarony{}

func init() {
	f := BLeitrim利特里姆.(*利特里姆LeitrimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leitrim",
		TitleName: "利特里姆",
		TitleCode: "b_leitrim",
	}
}
