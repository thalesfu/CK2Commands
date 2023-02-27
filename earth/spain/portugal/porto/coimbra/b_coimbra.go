package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科英布拉CoimbraBarony struct {
	feud.BaseBarony
}

var BCoimbra科英布拉 feud.Barony = &科英布拉CoimbraBarony{}

func init() {
    f := BCoimbra科英布拉.(*科英布拉CoimbraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coimbra",
		TitleName: "科英布拉",
		TitleCode: "b_coimbra",
	}
}
