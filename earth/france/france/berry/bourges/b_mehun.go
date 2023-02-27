package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默安MehunBarony struct {
	feud.BaseBarony
}

var BMehun默安 feud.Barony = &默安MehunBarony{}

func init() {
    f := BMehun默安.(*默安MehunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehun",
		TitleName: "默安",
		TitleCode: "b_mehun",
	}
}
