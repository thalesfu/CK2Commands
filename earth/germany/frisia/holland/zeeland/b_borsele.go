package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔瑟莱BorseleBarony struct {
	feud.BaseBarony
}

var BBorsele博尔瑟莱 feud.Barony = &博尔瑟莱BorseleBarony{}

func init() {
    f := BBorsele博尔瑟莱.(*博尔瑟莱BorseleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borsele",
		TitleName: "博尔瑟莱",
		TitleCode: "b_borsele",
	}
}
