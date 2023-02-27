package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳布卢斯NablusBarony struct {
	feud.BaseBarony
}

var BNablus纳布卢斯 feud.Barony = &纳布卢斯NablusBarony{}

func init() {
    f := BNablus纳布卢斯.(*纳布卢斯NablusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nablus",
		TitleName: "纳布卢斯",
		TitleCode: "b_nablus",
	}
}
