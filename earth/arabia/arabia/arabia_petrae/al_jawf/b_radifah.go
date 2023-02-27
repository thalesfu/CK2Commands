package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉迪法RadifahBarony struct {
	feud.BaseBarony
}

var BRadifah拉迪法 feud.Barony = &拉迪法RadifahBarony{}

func init() {
    f := BRadifah拉迪法.(*拉迪法RadifahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radifah",
		TitleName: "拉迪法",
		TitleCode: "b_radifah",
	}
}
