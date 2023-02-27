package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海鲁尔菲斯费尔德HaerulfisfeldBarony struct {
	feud.BaseBarony
}

var BHaerulfisfeld海鲁尔菲斯费尔德 feud.Barony = &海鲁尔菲斯费尔德HaerulfisfeldBarony{}

func init() {
    f := BHaerulfisfeld海鲁尔菲斯费尔德.(*海鲁尔菲斯费尔德HaerulfisfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haerulfisfeld",
		TitleName: "海鲁尔菲斯费尔德",
		TitleCode: "b_haerulfisfeld",
	}
}
