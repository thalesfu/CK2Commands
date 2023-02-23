package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基什KishBarony struct {
	feud.BaseBarony
}

var BKish基什 feud.Barony = &基什KishBarony{}

func init() {
	f := BKish基什.(*基什KishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kish",
		TitleName: "基什",
		TitleCode: "b_kish",
	}
}
