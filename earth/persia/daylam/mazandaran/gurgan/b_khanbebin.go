package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汗贝宾KhanbebinBarony struct {
	feud.BaseBarony
}

var BKhanbebin汗贝宾 feud.Barony = &汗贝宾KhanbebinBarony{}

func init() {
    f := BKhanbebin汗贝宾.(*汗贝宾KhanbebinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khanbebin",
		TitleName: "汗贝宾",
		TitleCode: "b_khanbebin",
	}
}
