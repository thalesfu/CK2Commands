package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多恩堡Dornburg_zeistBarony struct {
	feud.BaseBarony
}

var BDornburg_zeist多恩堡 feud.Barony = &多恩堡Dornburg_zeistBarony{}

func init() {
    f := BDornburg_zeist多恩堡.(*多恩堡Dornburg_zeistBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dornburg_zeist",
		TitleName: "多恩堡",
		TitleCode: "b_dornburg_zeist",
	}
}
