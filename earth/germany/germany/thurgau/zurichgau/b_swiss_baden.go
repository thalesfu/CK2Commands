package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴登Swiss_badenBarony struct {
	feud.BaseBarony
}

var BSwiss_baden巴登 feud.Barony = &巴登Swiss_badenBarony{}

func init() {
    f := BSwiss_baden巴登.(*巴登Swiss_badenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swiss_baden",
		TitleName: "巴登",
		TitleCode: "b_swiss_baden",
	}
}
