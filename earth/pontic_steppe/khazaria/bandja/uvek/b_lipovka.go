package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利波夫卡LipovkaBarony struct {
	feud.BaseBarony
}

var BLipovka利波夫卡 feud.Barony = &利波夫卡LipovkaBarony{}

func init() {
    f := BLipovka利波夫卡.(*利波夫卡LipovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lipovka",
		TitleName: "利波夫卡",
		TitleCode: "b_lipovka",
	}
}
