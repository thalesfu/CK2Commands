package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦代诺BondenoBarony struct {
	feud.BaseBarony
}

var BBondeno邦代诺 feud.Barony = &邦代诺BondenoBarony{}

func init() {
    f := BBondeno邦代诺.(*邦代诺BondenoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bondeno",
		TitleName: "邦代诺",
		TitleCode: "b_bondeno",
	}
}
