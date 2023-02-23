package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔日BourgesBarony struct {
	feud.BaseBarony
}

var BBourges布尔日 feud.Barony = &布尔日BourgesBarony{}

func init() {
	f := BBourges布尔日.(*布尔日BourgesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourges",
		TitleName: "布尔日",
		TitleCode: "b_bourges",
	}
}
