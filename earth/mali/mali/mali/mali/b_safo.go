package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨佛SafoBarony struct {
	feud.BaseBarony
}

var BSafo萨佛 feud.Barony = &萨佛SafoBarony{}

func init() {
    f := BSafo萨佛.(*萨佛SafoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safo",
		TitleName: "萨佛",
		TitleCode: "b_safo",
	}
}
