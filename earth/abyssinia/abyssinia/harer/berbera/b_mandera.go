package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼德拉ManderaBarony struct {
	feud.BaseBarony
}

var BMandera曼德拉 feud.Barony = &曼德拉ManderaBarony{}

func init() {
	f := BMandera曼德拉.(*曼德拉ManderaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandera",
		TitleName: "曼德拉",
		TitleCode: "b_mandera",
	}
}
