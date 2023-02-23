package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨武尔SavurBarony struct {
	feud.BaseBarony
}

var BSavur萨武尔 feud.Barony = &萨武尔SavurBarony{}

func init() {
	f := BSavur萨武尔.(*萨武尔SavurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savur",
		TitleName: "萨武尔",
		TitleCode: "b_savur",
	}
}
