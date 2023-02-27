package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里坤BalikunBarony struct {
	feud.BaseBarony
}

var BBalikun巴里坤 feud.Barony = &巴里坤BalikunBarony{}

func init() {
    f := BBalikun巴里坤.(*巴里坤BalikunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balikun",
		TitleName: "巴里坤",
		TitleCode: "b_balikun",
	}
}
