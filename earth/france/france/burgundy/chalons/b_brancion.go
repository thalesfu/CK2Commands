package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布朗雄BrancionBarony struct {
	feud.BaseBarony
}

var BBrancion布朗雄 feud.Barony = &布朗雄BrancionBarony{}

func init() {
	f := BBrancion布朗雄.(*布朗雄BrancionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brancion",
		TitleName: "布朗雄",
		TitleCode: "b_brancion",
	}
}
