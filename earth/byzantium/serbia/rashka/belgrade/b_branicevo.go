package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉尼切沃BranicevoBarony struct {
	feud.BaseBarony
}

var BBranicevo布拉尼切沃 feud.Barony = &布拉尼切沃BranicevoBarony{}

func init() {
	f := BBranicevo布拉尼切沃.(*布拉尼切沃BranicevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "branicevo",
		TitleName: "布拉尼切沃",
		TitleCode: "b_branicevo",
	}
}
