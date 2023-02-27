package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍沃HoveBarony struct {
	feud.BaseBarony
}

var BHove霍沃 feud.Barony = &霍沃HoveBarony{}

func init() {
    f := BHove霍沃.(*霍沃HoveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hove",
		TitleName: "霍沃",
		TitleCode: "b_hove",
	}
}
