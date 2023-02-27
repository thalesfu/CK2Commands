package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里茨拉尔FritzlarBarony struct {
	feud.BaseBarony
}

var BFritzlar弗里茨拉尔 feud.Barony = &弗里茨拉尔FritzlarBarony{}

func init() {
    f := BFritzlar弗里茨拉尔.(*弗里茨拉尔FritzlarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fritzlar",
		TitleName: "弗里茨拉尔",
		TitleCode: "b_fritzlar",
	}
}
