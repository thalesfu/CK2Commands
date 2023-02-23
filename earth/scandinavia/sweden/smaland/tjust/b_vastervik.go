package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯特维克VastervikBarony struct {
	feud.BaseBarony
}

var BVastervik韦斯特维克 feud.Barony = &韦斯特维克VastervikBarony{}

func init() {
	f := BVastervik韦斯特维克.(*韦斯特维克VastervikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vastervik",
		TitleName: "韦斯特维克",
		TitleCode: "b_vastervik",
	}
}
