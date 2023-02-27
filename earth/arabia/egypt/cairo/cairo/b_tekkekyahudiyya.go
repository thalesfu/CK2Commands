package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰勒阿呼蒂亚TekkekyahudiyyaBarony struct {
	feud.BaseBarony
}

var BTekkekyahudiyya泰勒阿呼蒂亚 feud.Barony = &泰勒阿呼蒂亚TekkekyahudiyyaBarony{}

func init() {
    f := BTekkekyahudiyya泰勒阿呼蒂亚.(*泰勒阿呼蒂亚TekkekyahudiyyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tekkekyahudiyya",
		TitleName: "泰勒阿呼蒂亚",
		TitleCode: "b_tekkekyahudiyya",
	}
}
