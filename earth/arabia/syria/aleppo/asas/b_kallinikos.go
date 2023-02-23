package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利尼克尤姆KallinikosBarony struct {
	feud.BaseBarony
}

var BKallinikos卡利尼克尤姆 feud.Barony = &卡利尼克尤姆KallinikosBarony{}

func init() {
	f := BKallinikos卡利尼克尤姆.(*卡利尼克尤姆KallinikosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kallinikos",
		TitleName: "卡利尼克尤姆",
		TitleCode: "b_kallinikos",
	}
}
