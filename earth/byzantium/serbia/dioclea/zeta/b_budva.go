package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布德瓦BudvaBarony struct {
	feud.BaseBarony
}

var BBudva布德瓦 feud.Barony = &布德瓦BudvaBarony{}

func init() {
    f := BBudva布德瓦.(*布德瓦BudvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "budva",
		TitleName: "布德瓦",
		TitleCode: "b_budva",
	}
}
