package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴贾BaajBarony struct {
	feud.BaseBarony
}

var BBaaj巴贾 feud.Barony = &巴贾BaajBarony{}

func init() {
	f := BBaaj巴贾.(*巴贾BaajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baaj",
		TitleName: "巴贾",
		TitleCode: "b_baaj",
	}
}
