package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉彻甘SalafcheganBarony struct {
	feud.BaseBarony
}

var BSalafchegan萨拉彻甘 feud.Barony = &萨拉彻甘SalafcheganBarony{}

func init() {
	f := BSalafchegan萨拉彻甘.(*萨拉彻甘SalafcheganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salafchegan",
		TitleName: "萨拉彻甘",
		TitleCode: "b_salafchegan",
	}
}
