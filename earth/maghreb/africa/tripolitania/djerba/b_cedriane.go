package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟德里安尼CedrianeBarony struct {
	feud.BaseBarony
}

var BCedriane瑟德里安尼 feud.Barony = &瑟德里安尼CedrianeBarony{}

func init() {
    f := BCedriane瑟德里安尼.(*瑟德里安尼CedrianeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cedriane",
		TitleName: "瑟德里安尼",
		TitleCode: "b_cedriane",
	}
}
