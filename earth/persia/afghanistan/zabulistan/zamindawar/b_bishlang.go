package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比什朗格BishlangBarony struct {
	feud.BaseBarony
}

var BBishlang比什朗格 feud.Barony = &比什朗格BishlangBarony{}

func init() {
    f := BBishlang比什朗格.(*比什朗格BishlangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bishlang",
		TitleName: "比什朗格",
		TitleCode: "b_bishlang",
	}
}
