package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁日BruggeBarony struct {
	feud.BaseBarony
}

var BBrugge布鲁日 feud.Barony = &布鲁日BruggeBarony{}

func init() {
	f := BBrugge布鲁日.(*布鲁日BruggeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brugge",
		TitleName: "布鲁日",
		TitleCode: "b_brugge",
	}
}
