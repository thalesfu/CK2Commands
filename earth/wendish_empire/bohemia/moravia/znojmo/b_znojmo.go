package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹诺伊莫ZnojmoBarony struct {
	feud.BaseBarony
}

var BZnojmo兹诺伊莫 feud.Barony = &兹诺伊莫ZnojmoBarony{}

func init() {
    f := BZnojmo兹诺伊莫.(*兹诺伊莫ZnojmoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "znojmo",
		TitleName: "兹诺伊莫",
		TitleCode: "b_znojmo",
	}
}
