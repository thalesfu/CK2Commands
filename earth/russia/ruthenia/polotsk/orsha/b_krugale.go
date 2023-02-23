package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁格韦KrugaleBarony struct {
	feud.BaseBarony
}

var BKrugale克鲁格韦 feud.Barony = &克鲁格韦KrugaleBarony{}

func init() {
	f := BKrugale克鲁格韦.(*克鲁格韦KrugaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krugale",
		TitleName: "克鲁格韦",
		TitleCode: "b_krugale",
	}
}
