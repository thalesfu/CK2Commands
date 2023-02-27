package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡坦扎罗CatanzaroBarony struct {
	feud.BaseBarony
}

var BCatanzaro卡坦扎罗 feud.Barony = &卡坦扎罗CatanzaroBarony{}

func init() {
    f := BCatanzaro卡坦扎罗.(*卡坦扎罗CatanzaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "catanzaro",
		TitleName: "卡坦扎罗",
		TitleCode: "b_catanzaro",
	}
}
