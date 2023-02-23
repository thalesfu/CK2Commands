package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别斯阔勒BeskolBarony struct {
	feud.BaseBarony
}

var BBeskol别斯阔勒 feud.Barony = &别斯阔勒BeskolBarony{}

func init() {
	f := BBeskol别斯阔勒.(*别斯阔勒BeskolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beskol",
		TitleName: "别斯阔勒",
		TitleCode: "b_beskol",
	}
}
