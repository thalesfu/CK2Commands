package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉格BuraqBarony struct {
	feud.BaseBarony
}

var BBuraq布拉格 feud.Barony = &布拉格BuraqBarony{}

func init() {
	f := BBuraq布拉格.(*布拉格BuraqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buraq",
		TitleName: "布拉格",
		TitleCode: "b_buraq",
	}
}
