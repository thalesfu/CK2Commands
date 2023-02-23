package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉贝德KalaberdBarony struct {
	feud.BaseBarony
}

var BKalaberd卡拉贝德 feud.Barony = &卡拉贝德KalaberdBarony{}

func init() {
	f := BKalaberd卡拉贝德.(*卡拉贝德KalaberdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalaberd",
		TitleName: "卡拉贝德",
		TitleCode: "b_kalaberd",
	}
}
