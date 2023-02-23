package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾托斯AetosBarony struct {
	feud.BaseBarony
}

var BAetos艾托斯 feud.Barony = &艾托斯AetosBarony{}

func init() {
	f := BAetos艾托斯.(*艾托斯AetosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aetos",
		TitleName: "艾托斯",
		TitleCode: "b_aetos",
	}
}
