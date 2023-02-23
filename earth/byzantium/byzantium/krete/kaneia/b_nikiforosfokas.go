package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼基佛洛佛卡斯NikiforosfokasBarony struct {
	feud.BaseBarony
}

var BNikiforosfokas尼基佛洛佛卡斯 feud.Barony = &尼基佛洛佛卡斯NikiforosfokasBarony{}

func init() {
	f := BNikiforosfokas尼基佛洛佛卡斯.(*尼基佛洛佛卡斯NikiforosfokasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikiforosfokas",
		TitleName: "尼基佛洛佛卡斯",
		TitleCode: "b_nikiforosfokas",
	}
}
