package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔加斯BourgasBarony struct {
	feud.BaseBarony
}

var BBourgas布尔加斯 feud.Barony = &布尔加斯BourgasBarony{}

func init() {
	f := BBourgas布尔加斯.(*布尔加斯BourgasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourgas",
		TitleName: "布尔加斯",
		TitleCode: "b_bourgas",
	}
}
