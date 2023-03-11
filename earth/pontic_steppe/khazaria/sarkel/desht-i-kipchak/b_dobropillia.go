package desht-i-kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多布罗皮利亚DobropilliaBarony struct {
	feud.BaseBarony
}

var BDobropillia多布罗皮利亚 feud.Barony = &多布罗皮利亚DobropilliaBarony{}

func init() {
    f := BDobropillia多布罗皮利亚.(*多布罗皮利亚DobropilliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobropillia",
		TitleName: "多布罗皮利亚",
		TitleCode: "b_dobropillia",
	}
}
