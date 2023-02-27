package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨勒马斯SalmasBarony struct {
	feud.BaseBarony
}

var BSalmas萨勒马斯 feud.Barony = &萨勒马斯SalmasBarony{}

func init() {
    f := BSalmas萨勒马斯.(*萨勒马斯SalmasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salmas",
		TitleName: "萨勒马斯",
		TitleCode: "b_salmas",
	}
}
