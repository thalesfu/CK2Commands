package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卡尔恰加SakarcageBarony struct {
	feud.BaseBarony
}

var BSakarcage萨卡尔恰加 feud.Barony = &萨卡尔恰加SakarcageBarony{}

func init() {
	f := BSakarcage萨卡尔恰加.(*萨卡尔恰加SakarcageBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakarcage",
		TitleName: "萨卡尔恰加",
		TitleCode: "b_sakarcage",
	}
}
