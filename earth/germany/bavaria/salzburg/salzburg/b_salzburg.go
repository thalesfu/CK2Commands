package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔茨堡SalzburgBarony struct {
	feud.BaseBarony
}

var BSalzburg萨尔茨堡 feud.Barony = &萨尔茨堡SalzburgBarony{}

func init() {
    f := BSalzburg萨尔茨堡.(*萨尔茨堡SalzburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salzburg",
		TitleName: "萨尔茨堡",
		TitleCode: "b_salzburg",
	}
}
