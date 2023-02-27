package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施塔姆斯StamsBarony struct {
	feud.BaseBarony
}

var BStams施塔姆斯 feud.Barony = &施塔姆斯StamsBarony{}

func init() {
    f := BStams施塔姆斯.(*施塔姆斯StamsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stams",
		TitleName: "施塔姆斯",
		TitleCode: "b_stams",
	}
}
