package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳雅克NajacBarony struct {
	feud.BaseBarony
}

var BNajac纳雅克 feud.Barony = &纳雅克NajacBarony{}

func init() {
    f := BNajac纳雅克.(*纳雅克NajacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "najac",
		TitleName: "纳雅克",
		TitleCode: "b_najac",
	}
}
