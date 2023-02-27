package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉扎佐LajazzoBarony struct {
	feud.BaseBarony
}

var BLajazzo拉扎佐 feud.Barony = &拉扎佐LajazzoBarony{}

func init() {
    f := BLajazzo拉扎佐.(*拉扎佐LajazzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lajazzo",
		TitleName: "拉扎佐",
		TitleCode: "b_lajazzo",
	}
}
