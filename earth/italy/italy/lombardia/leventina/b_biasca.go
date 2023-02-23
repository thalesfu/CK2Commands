package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比亚斯卡BiascaBarony struct {
	feud.BaseBarony
}

var BBiasca比亚斯卡 feud.Barony = &比亚斯卡BiascaBarony{}

func init() {
	f := BBiasca比亚斯卡.(*比亚斯卡BiascaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biasca",
		TitleName: "比亚斯卡",
		TitleCode: "b_biasca",
	}
}
