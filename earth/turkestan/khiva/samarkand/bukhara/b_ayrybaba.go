package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾里巴巴AyrybabaBarony struct {
	feud.BaseBarony
}

var BAyrybaba艾里巴巴 feud.Barony = &艾里巴巴AyrybabaBarony{}

func init() {
    f := BAyrybaba艾里巴巴.(*艾里巴巴AyrybabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayrybaba",
		TitleName: "艾里巴巴",
		TitleCode: "b_ayrybaba",
	}
}
