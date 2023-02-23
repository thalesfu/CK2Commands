package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗须卢BarasuruBarony struct {
	feud.BaseBarony
}

var BBarasuru婆罗须卢 feud.Barony = &婆罗须卢BarasuruBarony{}

func init() {
	f := BBarasuru婆罗须卢.(*婆罗须卢BarasuruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barasuru",
		TitleName: "婆罗须卢",
		TitleCode: "b_barasuru",
	}
}
