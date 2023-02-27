package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采尔布斯特ZerbstBarony struct {
	feud.BaseBarony
}

var BZerbst采尔布斯特 feud.Barony = &采尔布斯特ZerbstBarony{}

func init() {
    f := BZerbst采尔布斯特.(*采尔布斯特ZerbstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zerbst",
		TitleName: "采尔布斯特",
		TitleCode: "b_zerbst",
	}
}
