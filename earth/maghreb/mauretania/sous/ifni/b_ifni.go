package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊夫尼IfniBarony struct {
	feud.BaseBarony
}

var BIfni伊夫尼 feud.Barony = &伊夫尼IfniBarony{}

func init() {
	f := BIfni伊夫尼.(*伊夫尼IfniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ifni",
		TitleName: "伊夫尼",
		TitleCode: "b_ifni",
	}
}
