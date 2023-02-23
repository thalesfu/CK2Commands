package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨索斯ThasosBarony struct {
	feud.BaseBarony
}

var BThasos萨索斯 feud.Barony = &萨索斯ThasosBarony{}

func init() {
	f := BThasos萨索斯.(*萨索斯ThasosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thasos",
		TitleName: "萨索斯",
		TitleCode: "b_thasos",
	}
}
