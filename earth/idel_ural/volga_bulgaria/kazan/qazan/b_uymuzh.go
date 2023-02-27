package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维穆日UymuzhBarony struct {
	feud.BaseBarony
}

var BUymuzh维穆日 feud.Barony = &维穆日UymuzhBarony{}

func init() {
    f := BUymuzh维穆日.(*维穆日UymuzhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uymuzh",
		TitleName: "维穆日",
		TitleCode: "b_uymuzh",
	}
}
