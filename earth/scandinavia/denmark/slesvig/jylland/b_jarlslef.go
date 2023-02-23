package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶斯莱乌JarlslefBarony struct {
	feud.BaseBarony
}

var BJarlslef耶斯莱乌 feud.Barony = &耶斯莱乌JarlslefBarony{}

func init() {
	f := BJarlslef耶斯莱乌.(*耶斯莱乌JarlslefBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarlslef",
		TitleName: "耶斯莱乌",
		TitleCode: "b_jarlslef",
	}
}
