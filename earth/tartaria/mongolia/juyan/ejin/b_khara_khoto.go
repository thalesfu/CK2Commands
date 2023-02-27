package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑水城Khara_khotoBarony struct {
	feud.BaseBarony
}

var BKhara_khoto黑水城 feud.Barony = &黑水城Khara_khotoBarony{}

func init() {
    f := BKhara_khoto黑水城.(*黑水城Khara_khotoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khara_khoto",
		TitleName: "黑水城",
		TitleCode: "b_khara_khoto",
	}
}
