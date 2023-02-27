package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日尔曼尼科波利斯GermanicopolisBarony struct {
	feud.BaseBarony
}

var BGermanicopolis日尔曼尼科波利斯 feud.Barony = &日尔曼尼科波利斯GermanicopolisBarony{}

func init() {
    f := BGermanicopolis日尔曼尼科波利斯.(*日尔曼尼科波利斯GermanicopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "germanicopolis",
		TitleName: "日尔曼尼科波利斯",
		TitleCode: "b_germanicopolis",
	}
}
