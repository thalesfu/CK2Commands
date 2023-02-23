package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科洛尼亚KoloneiaBarony struct {
	feud.BaseBarony
}

var BKoloneia科洛尼亚 feud.Barony = &科洛尼亚KoloneiaBarony{}

func init() {
	f := BKoloneia科洛尼亚.(*科洛尼亚KoloneiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koloneia",
		TitleName: "科洛尼亚",
		TitleCode: "b_koloneia",
	}
}
