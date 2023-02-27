package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卢特卡SlutkaBarony struct {
	feud.BaseBarony
}

var BSlutka斯卢特卡 feud.Barony = &斯卢特卡SlutkaBarony{}

func init() {
    f := BSlutka斯卢特卡.(*斯卢特卡SlutkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slutka",
		TitleName: "斯卢特卡",
		TitleCode: "b_slutka",
	}
}
