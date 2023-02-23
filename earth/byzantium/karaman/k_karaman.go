package karaman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaramanKingdom interface {
	feud.Kingdom
}

type 卡拉曼KaramanKingdom struct {
	feud.BaseKingdom
}

var KKaraman卡拉曼 KaramanKingdom = &卡拉曼KaramanKingdom{}

func init() {
	f := KKaraman卡拉曼.(*卡拉曼KaramanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "karaman",
		TitleName: "卡拉曼",
		TitleCode: "k_karaman",
		Dukes:     map[string]feud.Duke{},
	}

}
