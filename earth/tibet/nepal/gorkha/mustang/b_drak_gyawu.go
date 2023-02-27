package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎甲吾Drak_gyawuBarony struct {
	feud.BaseBarony
}

var BDrak_gyawu扎甲吾 feud.Barony = &扎甲吾Drak_gyawuBarony{}

func init() {
    f := BDrak_gyawu扎甲吾.(*扎甲吾Drak_gyawuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drak_gyawu",
		TitleName: "扎甲吾",
		TitleCode: "b_drak_gyawu",
	}
}
