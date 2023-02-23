package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯特罗斯VasterasBarony struct {
	feud.BaseBarony
}

var BVasteras韦斯特罗斯 feud.Barony = &韦斯特罗斯VasterasBarony{}

func init() {
	f := BVasteras韦斯特罗斯.(*韦斯特罗斯VasterasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasteras",
		TitleName: "韦斯特罗斯",
		TitleCode: "b_vasteras",
	}
}
