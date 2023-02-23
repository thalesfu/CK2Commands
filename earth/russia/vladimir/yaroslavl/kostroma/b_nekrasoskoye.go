package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅克拉索斯科耶NekrasoskoyeBarony struct {
	feud.BaseBarony
}

var BNekrasoskoye涅克拉索斯科耶 feud.Barony = &涅克拉索斯科耶NekrasoskoyeBarony{}

func init() {
	f := BNekrasoskoye涅克拉索斯科耶.(*涅克拉索斯科耶NekrasoskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nekrasoskoye",
		TitleName: "涅克拉索斯科耶",
		TitleCode: "b_nekrasoskoye",
	}
}
