package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙圣万桑MontstvincentBarony struct {
	feud.BaseBarony
}

var BMontstvincent蒙圣万桑 feud.Barony = &蒙圣万桑MontstvincentBarony{}

func init() {
	f := BMontstvincent蒙圣万桑.(*蒙圣万桑MontstvincentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montstvincent",
		TitleName: "蒙圣万桑",
		TitleCode: "b_montstvincent",
	}
}
