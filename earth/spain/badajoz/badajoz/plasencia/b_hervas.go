package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔瓦斯HervasBarony struct {
	feud.BaseBarony
}

var BHervas埃尔瓦斯 feud.Barony = &埃尔瓦斯HervasBarony{}

func init() {
    f := BHervas埃尔瓦斯.(*埃尔瓦斯HervasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hervas",
		TitleName: "埃尔瓦斯",
		TitleCode: "b_hervas",
	}
}
