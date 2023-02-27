package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛古拉HingulaBarony struct {
	feud.BaseBarony
}

var BHingula辛古拉 feud.Barony = &辛古拉HingulaBarony{}

func init() {
    f := BHingula辛古拉.(*辛古拉HingulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hingula",
		TitleName: "辛古拉",
		TitleCode: "b_hingula",
	}
}
