package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶塔GetaBarony struct {
	feud.BaseBarony
}

var BGeta耶塔 feud.Barony = &耶塔GetaBarony{}

func init() {
    f := BGeta耶塔.(*耶塔GetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geta",
		TitleName: "耶塔",
		TitleCode: "b_geta",
	}
}
