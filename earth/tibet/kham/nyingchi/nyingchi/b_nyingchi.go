package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼池NyingchiBarony struct {
	feud.BaseBarony
}

var BNyingchi尼池 feud.Barony = &尼池NyingchiBarony{}

func init() {
    f := BNyingchi尼池.(*尼池NyingchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyingchi",
		TitleName: "尼池",
		TitleCode: "b_nyingchi",
	}
}
