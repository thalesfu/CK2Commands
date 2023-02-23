package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉加BragaBarony struct {
	feud.BaseBarony
}

var BBraga布拉加 feud.Barony = &布拉加BragaBarony{}

func init() {
	f := BBraga布拉加.(*布拉加BragaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "braga",
		TitleName: "布拉加",
		TitleCode: "b_braga",
	}
}
