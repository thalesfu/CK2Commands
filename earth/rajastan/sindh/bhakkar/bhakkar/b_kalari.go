package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉里KalariBarony struct {
	feud.BaseBarony
}

var BKalari卡拉里 feud.Barony = &卡拉里KalariBarony{}

func init() {
	f := BKalari卡拉里.(*卡拉里KalariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalari",
		TitleName: "卡拉里",
		TitleCode: "b_kalari",
	}
}
