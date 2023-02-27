package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦补罗他罗KaparthulaBarony struct {
	feud.BaseBarony
}

var BKaparthula迦补罗他罗 feud.Barony = &迦补罗他罗KaparthulaBarony{}

func init() {
    f := BKaparthula迦补罗他罗.(*迦补罗他罗KaparthulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaparthula",
		TitleName: "迦补罗他罗",
		TitleCode: "b_kaparthula",
	}
}
