package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八儿忽真BargujinBarony struct {
	feud.BaseBarony
}

var BBargujin八儿忽真 feud.Barony = &八儿忽真BargujinBarony{}

func init() {
	f := BBargujin八儿忽真.(*八儿忽真BargujinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bargujin",
		TitleName: "八儿忽真",
		TitleCode: "b_bargujin",
	}
}
