package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾达布尔AidabulBarony struct {
	feud.BaseBarony
}

var BAidabul艾达布尔 feud.Barony = &艾达布尔AidabulBarony{}

func init() {
	f := BAidabul艾达布尔.(*艾达布尔AidabulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aidabul",
		TitleName: "艾达布尔",
		TitleCode: "b_aidabul",
	}
}
