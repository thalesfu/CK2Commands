package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兹皮BzybBarony struct {
	feud.BaseBarony
}

var BBzyb布兹皮 feud.Barony = &布兹皮BzybBarony{}

func init() {
    f := BBzyb布兹皮.(*布兹皮BzybBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bzyb",
		TitleName: "布兹皮",
		TitleCode: "b_bzyb",
	}
}
