package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基兰达加迪Joranda_gadhiBarony struct {
	feud.BaseBarony
}

var BJoranda_gadhi基兰达加迪 feud.Barony = &基兰达加迪Joranda_gadhiBarony{}

func init() {
    f := BJoranda_gadhi基兰达加迪.(*基兰达加迪Joranda_gadhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joranda_gadhi",
		TitleName: "基兰达加迪",
		TitleCode: "b_joranda_gadhi",
	}
}
