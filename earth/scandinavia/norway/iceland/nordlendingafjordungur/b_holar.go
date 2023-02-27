package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍拉尔HolarBarony struct {
	feud.BaseBarony
}

var BHolar霍拉尔 feud.Barony = &霍拉尔HolarBarony{}

func init() {
    f := BHolar霍拉尔.(*霍拉尔HolarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holar",
		TitleName: "霍拉尔",
		TitleCode: "b_holar",
	}
}
