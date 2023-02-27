package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰帕莱TaipaleBarony struct {
	feud.BaseBarony
}

var BTaipale泰帕莱 feud.Barony = &泰帕莱TaipaleBarony{}

func init() {
    f := BTaipale泰帕莱.(*泰帕莱TaipaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taipale",
		TitleName: "泰帕莱",
		TitleCode: "b_taipale",
	}
}
