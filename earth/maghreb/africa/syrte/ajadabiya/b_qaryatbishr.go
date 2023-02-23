package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀鲁特比斯QaryatbishrBarony struct {
	feud.BaseBarony
}

var BQaryatbishr喀鲁特比斯 feud.Barony = &喀鲁特比斯QaryatbishrBarony{}

func init() {
	f := BQaryatbishr喀鲁特比斯.(*喀鲁特比斯QaryatbishrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaryatbishr",
		TitleName: "喀鲁特比斯",
		TitleCode: "b_qaryatbishr",
	}
}
