package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔索斯TarsosBarony struct {
	feud.BaseBarony
}

var BTarsos塔索斯 feud.Barony = &塔索斯TarsosBarony{}

func init() {
    f := BTarsos塔索斯.(*塔索斯TarsosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarsos",
		TitleName: "塔索斯",
		TitleCode: "b_tarsos",
	}
}
