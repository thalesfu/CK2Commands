package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆斯蒂里El_mstiriBarony struct {
	feud.BaseBarony
}

var BEl_mstiri姆斯蒂里 feud.Barony = &姆斯蒂里El_mstiriBarony{}

func init() {
    f := BEl_mstiri姆斯蒂里.(*姆斯蒂里El_mstiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_mstiri",
		TitleName: "姆斯蒂里",
		TitleCode: "b_el_mstiri",
	}
}
