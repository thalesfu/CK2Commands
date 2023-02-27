package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基多尼亚KydoniaBarony struct {
	feud.BaseBarony
}

var BKydonia基多尼亚 feud.Barony = &基多尼亚KydoniaBarony{}

func init() {
    f := BKydonia基多尼亚.(*基多尼亚KydoniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kydonia",
		TitleName: "基多尼亚",
		TitleCode: "b_kydonia",
	}
}
