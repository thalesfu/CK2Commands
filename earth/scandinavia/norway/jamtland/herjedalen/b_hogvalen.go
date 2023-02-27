package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫格沃伦HogvalenBarony struct {
	feud.BaseBarony
}

var BHogvalen赫格沃伦 feud.Barony = &赫格沃伦HogvalenBarony{}

func init() {
    f := BHogvalen赫格沃伦.(*赫格沃伦HogvalenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hogvalen",
		TitleName: "赫格沃伦",
		TitleCode: "b_hogvalen",
	}
}
