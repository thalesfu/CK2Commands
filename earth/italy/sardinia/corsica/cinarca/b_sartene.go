package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔泰讷SarteneBarony struct {
	feud.BaseBarony
}

var BSartene萨尔泰讷 feud.Barony = &萨尔泰讷SarteneBarony{}

func init() {
	f := BSartene萨尔泰讷.(*萨尔泰讷SarteneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sartene",
		TitleName: "萨尔泰讷",
		TitleCode: "b_sartene",
	}
}
