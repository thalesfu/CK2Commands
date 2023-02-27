package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊夫雷亚IvreaBarony struct {
	feud.BaseBarony
}

var BIvrea伊夫雷亚 feud.Barony = &伊夫雷亚IvreaBarony{}

func init() {
    f := BIvrea伊夫雷亚.(*伊夫雷亚IvreaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivrea",
		TitleName: "伊夫雷亚",
		TitleCode: "b_ivrea",
	}
}
