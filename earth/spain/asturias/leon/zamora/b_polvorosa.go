package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔沃罗萨PolvorosaBarony struct {
	feud.BaseBarony
}

var BPolvorosa波尔沃罗萨 feud.Barony = &波尔沃罗萨PolvorosaBarony{}

func init() {
    f := BPolvorosa波尔沃罗萨.(*波尔沃罗萨PolvorosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polvorosa",
		TitleName: "波尔沃罗萨",
		TitleCode: "b_polvorosa",
	}
}
