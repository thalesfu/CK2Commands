package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁平RuppinBarony struct {
	feud.BaseBarony
}

var BRuppin鲁平 feud.Barony = &鲁平RuppinBarony{}

func init() {
    f := BRuppin鲁平.(*鲁平RuppinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruppin",
		TitleName: "鲁平",
		TitleCode: "b_ruppin",
	}
}
