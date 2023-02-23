package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提格利尼亚TigrinyaBarony struct {
	feud.BaseBarony
}

var BTigrinya提格利尼亚 feud.Barony = &提格利尼亚TigrinyaBarony{}

func init() {
	f := BTigrinya提格利尼亚.(*提格利尼亚TigrinyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigrinya",
		TitleName: "提格利尼亚",
		TitleCode: "b_tigrinya",
	}
}
