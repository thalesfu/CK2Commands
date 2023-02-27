package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旺多姆VendomeBarony struct {
	feud.BaseBarony
}

var BVendome旺多姆 feud.Barony = &旺多姆VendomeBarony{}

func init() {
    f := BVendome旺多姆.(*旺多姆VendomeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vendome",
		TitleName: "旺多姆",
		TitleCode: "b_vendome",
	}
}
