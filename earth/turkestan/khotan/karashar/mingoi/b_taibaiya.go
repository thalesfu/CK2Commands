package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 太白垭TaibaiyaBarony struct {
	feud.BaseBarony
}

var BTaibaiya太白垭 feud.Barony = &太白垭TaibaiyaBarony{}

func init() {
    f := BTaibaiya太白垭.(*太白垭TaibaiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taibaiya",
		TitleName: "太白垭",
		TitleCode: "b_taibaiya",
	}
}
