package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡卡尔堡AlcaladeljucarBarony struct {
	feud.BaseBarony
}

var BAlcaladeljucar胡卡尔堡 feud.Barony = &胡卡尔堡AlcaladeljucarBarony{}

func init() {
    f := BAlcaladeljucar胡卡尔堡.(*胡卡尔堡AlcaladeljucarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcaladeljucar",
		TitleName: "胡卡尔堡",
		TitleCode: "b_alcaladeljucar",
	}
}
