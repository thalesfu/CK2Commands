package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布迪姆利亚BudimljaBarony struct {
	feud.BaseBarony
}

var BBudimlja布迪姆利亚 feud.Barony = &布迪姆利亚BudimljaBarony{}

func init() {
	f := BBudimlja布迪姆利亚.(*布迪姆利亚BudimljaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "budimlja",
		TitleName: "布迪姆利亚",
		TitleCode: "b_budimlja",
	}
}
