package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨翁林纳SavonlinnaBarony struct {
	feud.BaseBarony
}

var BSavonlinna萨翁林纳 feud.Barony = &萨翁林纳SavonlinnaBarony{}

func init() {
	f := BSavonlinna萨翁林纳.(*萨翁林纳SavonlinnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savonlinna",
		TitleName: "萨翁林纳",
		TitleCode: "b_savonlinna",
	}
}
