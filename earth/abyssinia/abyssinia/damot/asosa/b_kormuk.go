package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔穆克KormukBarony struct {
	feud.BaseBarony
}

var BKormuk库尔穆克 feud.Barony = &库尔穆克KormukBarony{}

func init() {
	f := BKormuk库尔穆克.(*库尔穆克KormukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kormuk",
		TitleName: "库尔穆克",
		TitleCode: "b_kormuk",
	}
}
