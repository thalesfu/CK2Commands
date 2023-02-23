package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼扎NizzaBarony struct {
	feud.BaseBarony
}

var BNizza尼扎 feud.Barony = &尼扎NizzaBarony{}

func init() {
	f := BNizza尼扎.(*尼扎NizzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizza",
		TitleName: "尼扎",
		TitleCode: "b_nizza",
	}
}
