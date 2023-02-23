package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因舍伊尔AinechchairBarony struct {
	feud.BaseBarony
}

var BAinechchair艾因舍伊尔 feud.Barony = &艾因舍伊尔AinechchairBarony{}

func init() {
	f := BAinechchair艾因舍伊尔.(*艾因舍伊尔AinechchairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainechchair",
		TitleName: "艾因舍伊尔",
		TitleCode: "b_ainechchair",
	}
}
