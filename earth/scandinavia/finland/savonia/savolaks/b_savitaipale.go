package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨维泰帕莱SavitaipaleBarony struct {
	feud.BaseBarony
}

var BSavitaipale萨维泰帕莱 feud.Barony = &萨维泰帕莱SavitaipaleBarony{}

func init() {
	f := BSavitaipale萨维泰帕莱.(*萨维泰帕莱SavitaipaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savitaipale",
		TitleName: "萨维泰帕莱",
		TitleCode: "b_savitaipale",
	}
}
