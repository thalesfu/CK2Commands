package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆赖萨斯Umm_ar_rasasBarony struct {
	feud.BaseBarony
}

var BUmm_ar_rasas乌姆赖萨斯 feud.Barony = &乌姆赖萨斯Umm_ar_rasasBarony{}

func init() {
    f := BUmm_ar_rasas乌姆赖萨斯.(*乌姆赖萨斯Umm_ar_rasasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umm_ar_rasas",
		TitleName: "乌姆赖萨斯",
		TitleCode: "b_umm_ar-rasas",
	}
}
