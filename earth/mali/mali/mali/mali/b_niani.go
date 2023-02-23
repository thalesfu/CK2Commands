package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼亚尼NianiBarony struct {
	feud.BaseBarony
}

var BNiani尼亚尼 feud.Barony = &尼亚尼NianiBarony{}

func init() {
	f := BNiani尼亚尼.(*尼亚尼NianiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niani",
		TitleName: "尼亚尼",
		TitleCode: "b_niani",
	}
}
