package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋波罗VatpatrakBarony struct {
	feud.BaseBarony
}

var BVatpatrak跋波罗 feud.Barony = &跋波罗VatpatrakBarony{}

func init() {
    f := BVatpatrak跋波罗.(*跋波罗VatpatrakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vatpatrak",
		TitleName: "跋波罗",
		TitleCode: "b_vatpatrak",
	}
}
