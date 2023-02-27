package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马德米尼亚MadminiyaBarony struct {
	feud.BaseBarony
}

var BMadminiya马德米尼亚 feud.Barony = &马德米尼亚MadminiyaBarony{}

func init() {
    f := BMadminiya马德米尼亚.(*马德米尼亚MadminiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madminiya",
		TitleName: "马德米尼亚",
		TitleCode: "b_madminiya",
	}
}
