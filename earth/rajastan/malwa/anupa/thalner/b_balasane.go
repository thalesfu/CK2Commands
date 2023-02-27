package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗娑内BalasaneBarony struct {
	feud.BaseBarony
}

var BBalasane婆罗娑内 feud.Barony = &婆罗娑内BalasaneBarony{}

func init() {
    f := BBalasane婆罗娑内.(*婆罗娑内BalasaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balasane",
		TitleName: "婆罗娑内",
		TitleCode: "b_balasane",
	}
}
