package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰希尔TaishirBarony struct {
	feud.BaseBarony
}

var BTaishir泰希尔 feud.Barony = &泰希尔TaishirBarony{}

func init() {
	f := BTaishir泰希尔.(*泰希尔TaishirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taishir",
		TitleName: "泰希尔",
		TitleCode: "b_taishir",
	}
}
