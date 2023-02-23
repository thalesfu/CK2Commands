package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布德诺BudnoBarony struct {
	feud.BaseBarony
}

var BBudno布德诺 feud.Barony = &布德诺BudnoBarony{}

func init() {
	f := BBudno布德诺.(*布德诺BudnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "budno",
		TitleName: "布德诺",
		TitleCode: "b_budno",
	}
}
