package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉赫蒂LahtiBarony struct {
	feud.BaseBarony
}

var BLahti拉赫蒂 feud.Barony = &拉赫蒂LahtiBarony{}

func init() {
	f := BLahti拉赫蒂.(*拉赫蒂LahtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahti",
		TitleName: "拉赫蒂",
		TitleCode: "b_lahti",
	}
}
