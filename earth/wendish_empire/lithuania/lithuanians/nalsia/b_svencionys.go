package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什文乔尼斯SvencionysBarony struct {
	feud.BaseBarony
}

var BSvencionys什文乔尼斯 feud.Barony = &什文乔尼斯SvencionysBarony{}

func init() {
    f := BSvencionys什文乔尼斯.(*什文乔尼斯SvencionysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svencionys",
		TitleName: "什文乔尼斯",
		TitleCode: "b_svencionys",
	}
}
