package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴卜塔宰BabtazaBarony struct {
	feud.BaseBarony
}

var BBabtaza巴卜塔宰 feud.Barony = &巴卜塔宰BabtazaBarony{}

func init() {
    f := BBabtaza巴卜塔宰.(*巴卜塔宰BabtazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babtaza",
		TitleName: "巴卜塔宰",
		TitleCode: "b_babtaza",
	}
}
