package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗尔巴赫RohrbachBarony struct {
	feud.BaseBarony
}

var BRohrbach罗尔巴赫 feud.Barony = &罗尔巴赫RohrbachBarony{}

func init() {
	f := BRohrbach罗尔巴赫.(*罗尔巴赫RohrbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohrbach",
		TitleName: "罗尔巴赫",
		TitleCode: "b_rohrbach",
	}
}
