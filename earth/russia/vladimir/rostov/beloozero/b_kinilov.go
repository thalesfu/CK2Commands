package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尼洛夫KinilovBarony struct {
	feud.BaseBarony
}

var BKinilov基尼洛夫 feud.Barony = &基尼洛夫KinilovBarony{}

func init() {
    f := BKinilov基尼洛夫.(*基尼洛夫KinilovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinilov",
		TitleName: "基尼洛夫",
		TitleCode: "b_kinilov",
	}
}
