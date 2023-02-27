package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔萨诺夫KirsanovBarony struct {
	feud.BaseBarony
}

var BKirsanov基尔萨诺夫 feud.Barony = &基尔萨诺夫KirsanovBarony{}

func init() {
    f := BKirsanov基尔萨诺夫.(*基尔萨诺夫KirsanovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirsanov",
		TitleName: "基尔萨诺夫",
		TitleCode: "b_kirsanov",
	}
}
