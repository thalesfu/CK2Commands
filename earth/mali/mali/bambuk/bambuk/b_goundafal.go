package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡达法尔GoundafalBarony struct {
	feud.BaseBarony
}

var BGoundafal贡达法尔 feud.Barony = &贡达法尔GoundafalBarony{}

func init() {
	f := BGoundafal贡达法尔.(*贡达法尔GoundafalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goundafal",
		TitleName: "贡达法尔",
		TitleCode: "b_goundafal",
	}
}
