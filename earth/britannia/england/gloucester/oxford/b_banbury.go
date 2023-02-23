package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班伯里BanburyBarony struct {
	feud.BaseBarony
}

var BBanbury班伯里 feud.Barony = &班伯里BanburyBarony{}

func init() {
	f := BBanbury班伯里.(*班伯里BanburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banbury",
		TitleName: "班伯里",
		TitleCode: "b_banbury",
	}
}
