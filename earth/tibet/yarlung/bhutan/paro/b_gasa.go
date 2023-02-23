package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加萨GasaBarony struct {
	feud.BaseBarony
}

var BGasa加萨 feud.Barony = &加萨GasaBarony{}

func init() {
	f := BGasa加萨.(*加萨GasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gasa",
		TitleName: "加萨",
		TitleCode: "b_gasa",
	}
}
