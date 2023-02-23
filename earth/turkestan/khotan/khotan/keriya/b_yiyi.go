package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊夷YiyiBarony struct {
	feud.BaseBarony
}

var BYiyi伊夷 feud.Barony = &伊夷YiyiBarony{}

func init() {
	f := BYiyi伊夷.(*伊夷YiyiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yiyi",
		TitleName: "伊夷",
		TitleCode: "b_yiyi",
	}
}
