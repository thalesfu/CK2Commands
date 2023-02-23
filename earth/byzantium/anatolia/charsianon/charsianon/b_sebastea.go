package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞巴斯蒂亚SebasteaBarony struct {
	feud.BaseBarony
}

var BSebastea塞巴斯蒂亚 feud.Barony = &塞巴斯蒂亚SebasteaBarony{}

func init() {
	f := BSebastea塞巴斯蒂亚.(*塞巴斯蒂亚SebasteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebastea",
		TitleName: "塞巴斯蒂亚",
		TitleCode: "b_sebastea",
	}
}
