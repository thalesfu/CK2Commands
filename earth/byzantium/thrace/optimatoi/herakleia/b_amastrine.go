package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马斯特里尼AmastrineBarony struct {
	feud.BaseBarony
}

var BAmastrine阿马斯特里尼 feud.Barony = &阿马斯特里尼AmastrineBarony{}

func init() {
    f := BAmastrine阿马斯特里尼.(*阿马斯特里尼AmastrineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amastrine",
		TitleName: "阿马斯特里尼",
		TitleCode: "b_amastrine",
	}
}
