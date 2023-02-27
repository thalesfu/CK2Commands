package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑楼迦尼SaruganiBarony struct {
	feud.BaseBarony
}

var BSarugani娑楼迦尼 feud.Barony = &娑楼迦尼SaruganiBarony{}

func init() {
    f := BSarugani娑楼迦尼.(*娑楼迦尼SaruganiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarugani",
		TitleName: "娑楼迦尼",
		TitleCode: "b_sarugani",
	}
}
