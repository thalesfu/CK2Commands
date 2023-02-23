package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特里尼齐KastrinitsiBarony struct {
	feud.BaseBarony
}

var BKastrinitsi卡斯特里尼齐 feud.Barony = &卡斯特里尼齐KastrinitsiBarony{}

func init() {
	f := BKastrinitsi卡斯特里尼齐.(*卡斯特里尼齐KastrinitsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastrinitsi",
		TitleName: "卡斯特里尼齐",
		TitleCode: "b_kastrinitsi",
	}
}
