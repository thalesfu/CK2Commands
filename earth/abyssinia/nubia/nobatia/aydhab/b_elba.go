package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾尔巴ElbaBarony struct {
	feud.BaseBarony
}

var BElba艾尔巴 feud.Barony = &艾尔巴ElbaBarony{}

func init() {
    f := BElba艾尔巴.(*艾尔巴ElbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elba",
		TitleName: "艾尔巴",
		TitleCode: "b_elba",
	}
}
