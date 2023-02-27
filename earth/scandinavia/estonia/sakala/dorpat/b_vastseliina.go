package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦斯采利纳VastseliinaBarony struct {
	feud.BaseBarony
}

var BVastseliina瓦斯采利纳 feud.Barony = &瓦斯采利纳VastseliinaBarony{}

func init() {
    f := BVastseliina瓦斯采利纳.(*瓦斯采利纳VastseliinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vastseliina",
		TitleName: "瓦斯采利纳",
		TitleCode: "b_vastseliina",
	}
}
