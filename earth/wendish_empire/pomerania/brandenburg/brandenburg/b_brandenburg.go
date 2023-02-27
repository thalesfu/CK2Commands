package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勃兰登堡BrandenburgBarony struct {
	feud.BaseBarony
}

var BBrandenburg勃兰登堡 feud.Barony = &勃兰登堡BrandenburgBarony{}

func init() {
    f := BBrandenburg勃兰登堡.(*勃兰登堡BrandenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brandenburg",
		TitleName: "勃兰登堡",
		TitleCode: "b_brandenburg",
	}
}
