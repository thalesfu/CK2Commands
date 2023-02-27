package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏梅SumyBarony struct {
	feud.BaseBarony
}

var BSumy苏梅 feud.Barony = &苏梅SumyBarony{}

func init() {
    f := BSumy苏梅.(*苏梅SumyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumy",
		TitleName: "苏梅",
		TitleCode: "b_sumy",
	}
}
