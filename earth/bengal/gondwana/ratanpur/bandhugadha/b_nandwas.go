package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难瓦斯NandwasBarony struct {
	feud.BaseBarony
}

var BNandwas难瓦斯 feud.Barony = &难瓦斯NandwasBarony{}

func init() {
    f := BNandwas难瓦斯.(*难瓦斯NandwasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandwas",
		TitleName: "难瓦斯",
		TitleCode: "b_nandwas",
	}
}
