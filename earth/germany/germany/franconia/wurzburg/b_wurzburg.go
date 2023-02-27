package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔茨堡WurzburgBarony struct {
	feud.BaseBarony
}

var BWurzburg维尔茨堡 feud.Barony = &维尔茨堡WurzburgBarony{}

func init() {
    f := BWurzburg维尔茨堡.(*维尔茨堡WurzburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wurzburg",
		TitleName: "维尔茨堡",
		TitleCode: "b_wurzburg",
	}
}
