package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉堡HamburgBarony struct {
	feud.BaseBarony
}

var BHamburg汉堡 feud.Barony = &汉堡HamburgBarony{}

func init() {
    f := BHamburg汉堡.(*汉堡HamburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamburg",
		TitleName: "汉堡",
		TitleCode: "b_hamburg",
	}
}
