package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼永NyonsBarony struct {
	feud.BaseBarony
}

var BNyons尼永 feud.Barony = &尼永NyonsBarony{}

func init() {
	f := BNyons尼永.(*尼永NyonsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyons",
		TitleName: "尼永",
		TitleCode: "b_nyons",
	}
}
