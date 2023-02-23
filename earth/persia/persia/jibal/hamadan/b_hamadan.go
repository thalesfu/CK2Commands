package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马丹HamadanBarony struct {
	feud.BaseBarony
}

var BHamadan哈马丹 feud.Barony = &哈马丹HamadanBarony{}

func init() {
	f := BHamadan哈马丹.(*哈马丹HamadanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamadan",
		TitleName: "哈马丹",
		TitleCode: "b_hamadan",
	}
}
