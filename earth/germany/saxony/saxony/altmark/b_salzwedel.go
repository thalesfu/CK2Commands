package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔茨韦德尔SalzwedelBarony struct {
	feud.BaseBarony
}

var BSalzwedel萨尔茨韦德尔 feud.Barony = &萨尔茨韦德尔SalzwedelBarony{}

func init() {
    f := BSalzwedel萨尔茨韦德尔.(*萨尔茨韦德尔SalzwedelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salzwedel",
		TitleName: "萨尔茨韦德尔",
		TitleCode: "b_salzwedel",
	}
}
