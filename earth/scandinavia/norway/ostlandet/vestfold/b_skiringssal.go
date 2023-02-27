package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯基林格斯萨SkiringssalBarony struct {
	feud.BaseBarony
}

var BSkiringssal斯基林格斯萨 feud.Barony = &斯基林格斯萨SkiringssalBarony{}

func init() {
    f := BSkiringssal斯基林格斯萨.(*斯基林格斯萨SkiringssalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skiringssal",
		TitleName: "斯基林格斯萨",
		TitleCode: "b_skiringssal",
	}
}
