package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗赖辛FreisingBarony struct {
	feud.BaseBarony
}

var BFreising弗赖辛 feud.Barony = &弗赖辛FreisingBarony{}

func init() {
	f := BFreising弗赖辛.(*弗赖辛FreisingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "freising",
		TitleName: "弗赖辛",
		TitleCode: "b_freising",
	}
}
