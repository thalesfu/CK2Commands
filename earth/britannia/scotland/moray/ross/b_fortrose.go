package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福特罗斯FortroseBarony struct {
	feud.BaseBarony
}

var BFortrose福特罗斯 feud.Barony = &福特罗斯FortroseBarony{}

func init() {
	f := BFortrose福特罗斯.(*福特罗斯FortroseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fortrose",
		TitleName: "福特罗斯",
		TitleCode: "b_fortrose",
	}
}
