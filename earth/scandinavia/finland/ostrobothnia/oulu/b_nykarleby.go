package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新卡勒比NykarlebyBarony struct {
	feud.BaseBarony
}

var BNykarleby新卡勒比 feud.Barony = &新卡勒比NykarlebyBarony{}

func init() {
    f := BNykarleby新卡勒比.(*新卡勒比NykarlebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nykarleby",
		TitleName: "新卡勒比",
		TitleCode: "b_nykarleby",
	}
}
