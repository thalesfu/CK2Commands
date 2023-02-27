package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科迦KhorzhakBarony struct {
	feud.BaseBarony
}

var BKhorzhak科迦 feud.Barony = &科迦KhorzhakBarony{}

func init() {
    f := BKhorzhak科迦.(*科迦KhorzhakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorzhak",
		TitleName: "科迦",
		TitleCode: "b_khorzhak",
	}
}
