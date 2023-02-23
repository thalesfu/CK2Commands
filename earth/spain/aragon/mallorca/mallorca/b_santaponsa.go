package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣蓬萨SantaponsaBarony struct {
	feud.BaseBarony
}

var BSantaponsa圣蓬萨 feud.Barony = &圣蓬萨SantaponsaBarony{}

func init() {
	f := BSantaponsa圣蓬萨.(*圣蓬萨SantaponsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santaponsa",
		TitleName: "圣蓬萨",
		TitleCode: "b_santaponsa",
	}
}
