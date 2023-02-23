package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯特胡斯VasterhusBarony struct {
	feud.BaseBarony
}

var BVasterhus韦斯特胡斯 feud.Barony = &韦斯特胡斯VasterhusBarony{}

func init() {
	f := BVasterhus韦斯特胡斯.(*韦斯特胡斯VasterhusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasterhus",
		TitleName: "韦斯特胡斯",
		TitleCode: "b_vasterhus",
	}
}
