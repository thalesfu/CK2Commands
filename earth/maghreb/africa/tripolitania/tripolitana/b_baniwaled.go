package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜尼沃利德BaniwaledBarony struct {
	feud.BaseBarony
}

var BBaniwaled拜尼沃利德 feud.Barony = &拜尼沃利德BaniwaledBarony{}

func init() {
	f := BBaniwaled拜尼沃利德.(*拜尼沃利德BaniwaledBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baniwaled",
		TitleName: "拜尼沃利德",
		TitleCode: "b_baniwaled",
	}
}
