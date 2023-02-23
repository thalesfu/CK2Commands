package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维耶尔宗VierzonBarony struct {
	feud.BaseBarony
}

var BVierzon维耶尔宗 feud.Barony = &维耶尔宗VierzonBarony{}

func init() {
	f := BVierzon维耶尔宗.(*维耶尔宗VierzonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vierzon",
		TitleName: "维耶尔宗",
		TitleCode: "b_vierzon",
	}
}
