package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卡尔SackalBarony struct {
	feud.BaseBarony
}

var BSackal萨卡尔 feud.Barony = &萨卡尔SackalBarony{}

func init() {
    f := BSackal萨卡尔.(*萨卡尔SackalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sackal",
		TitleName: "萨卡尔",
		TitleCode: "b_sackal",
	}
}
