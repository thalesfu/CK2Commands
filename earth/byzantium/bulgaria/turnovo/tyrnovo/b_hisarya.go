package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希萨里亚HisaryaBarony struct {
	feud.BaseBarony
}

var BHisarya希萨里亚 feud.Barony = &希萨里亚HisaryaBarony{}

func init() {
    f := BHisarya希萨里亚.(*希萨里亚HisaryaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hisarya",
		TitleName: "希萨里亚",
		TitleCode: "b_hisarya",
	}
}
